package ws

import (
	"context"
	"database/sql"
	constant "ginchat2/common"
	"ginchat2/internal/kafka"
	"ginchat2/internal/svc"
	"ginchat2/models"
	"ginchat2/protocol"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

type inboundMsg struct {
	client *Client
	data   []byte
}

// Hub tracks websocket clients and distributes messages.
type Hub struct {
	Register   chan *Client
	Unregister chan *Client
	Inbound    chan inboundMsg
	BusIn      chan []byte
	Clients    map[int64]map[*Client]struct{}
	svcCtx     *svc.ServiceContext
}

// NewHub constructs a Hub and starts its run loop.
func NewHub(svcCtx *svc.ServiceContext) *Hub {
	h := &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Inbound:    make(chan inboundMsg, 1024),
		BusIn:      make(chan []byte, 1024),
		Clients:    make(map[int64]map[*Client]struct{}),
		svcCtx:     svcCtx,
	}
	go h.run()
	if svcCtx != nil && svcCtx.Config.MsgChannel.ChannelType == constant.KAFKA {
		go kafka.ConsumerMsg(func(data []byte) {
			h.BusIn <- data
		})
	}
	return h
}

func (h *Hub) run() {
	logx.Info("staring server")
	for {
		select {
		case conn := <-h.Register:
			logx.Infof("new user login in %v", conn.UserID)
			if h.Clients[conn.UserID] == nil {
				h.Clients[conn.UserID] = make(map[*Client]struct{})
			}
			h.Clients[conn.UserID][conn] = struct{}{}
			msg := &protocol.Message{
				From:         0,
				FromUsername: "System",
				To:           conn.UserID,
				Content:      "welcome!",
			}
			protoMsg, _ := proto.Marshal(msg)
			conn.Send <- protoMsg

		case conn := <-h.Unregister:
			logx.Infof("logout %v", conn.UserID)
			if deviceSet, ok := h.Clients[conn.UserID]; ok {
				delete(deviceSet, conn)
				if len(deviceSet) == 0 {
					delete(h.Clients, conn.UserID)
				}
			}

		case in := <-h.Inbound:
			h.handleInbound(in)

		case message := <-h.BusIn:
			h.handleBusMessage(message)
		}
	}
}

func (h *Hub) handleInbound(in inboundMsg) {
	msg := &protocol.Message{}
	if err := proto.Unmarshal(in.data, msg); err != nil {
		logx.Error("Unmarshal fail!:", err)
		return
	}
	// 防止客户端伪造发送者
	msg.From = in.client.UserID

	// 这里先按最小可用：只存文本/音视频等(1..5)；文件/图片二进制暂不落库
	if msg.ContentType == constant.TEXT {
		persona, err := h.svcCtx.PersonaModel.FindOne(context.Background(), msg.Persona)
		if err != nil {
			logx.Infof("%v:%v", in.client.UserID, err)
			return
		}
		if persona.UserId != in.client.UserID {
			logx.Infof("%v:use %v not own him.", in.client.UserID, persona.Id)
		}
		result, err := h.svcCtx.Gateway.Process(in.client.ctx, persona.Prompt, msg.Content, "")
		if err != nil {
			logx.Error(err)
			return
		}
		msg.Content = result
	}
	if msg.To != 0 && msg.ContentType >= constant.TEXT && msg.ContentType <= constant.VIDEO {
		_, err := h.svcCtx.MessageModel.Insert(context.Background(), &models.Message{
			FromUserId:  sql.NullInt64{Int64: msg.From, Valid: true},
			ToUserId:    sql.NullInt64{Int64: msg.To, Valid: true},
			Content:     msg.Content,
			Url:         msg.Url,
			MessageType: sql.NullInt64{Int64: int64(msg.MessageType), Valid: true},
			ContentType: sql.NullInt64{Int64: int64(msg.ContentType), Valid: true},
		})
		if err != nil {
			logx.Error("save message fail:", err)
		}
	}

	out, err := proto.Marshal(msg)
	if err != nil {
		logx.Error("Marshal fail:", err)
		return
	}
	if h.svcCtx.Config.MsgChannel.ChannelType == constant.KAFKA {
		kafka.Send(out)
		return
	}
	h.deliver(out)
}

func (h *Hub) handleBusMessage(message []byte) {
	// 来自 Kafka 的消息：只负责投递，不重复落库
	h.deliver(message)
}

func (h *Hub) deliver(message []byte) {
	msg := &protocol.Message{}
	if err := proto.Unmarshal(message, msg); err != nil {
		logx.Error("Unmarshal fail!:", err)
		return
	}
	if msg.To != 0 {
		switch msg.MessageType {
		case constant.MESSAGE_TYPE_USER:
			h.sendToUser(msg.To, message)
		case constant.MESSAGE_TYPE_GROUP:
			// TODO: group message
		}
		return
	}

	for _, deviceSet := range h.Clients {
		for conn := range deviceSet {
			select {
			case conn.Send <- message:
			default:
				delete(deviceSet, conn)
				close(conn.Send)
			}
		}
	}
}

func (h *Hub) sendToUser(userID int64, message []byte) {
	deviceSet, ok := h.Clients[userID]
	if !ok {
		logx.Errorf("Clients dont have %v", userID)
		return
	}
	for conn := range deviceSet {
		select {
		case conn.Send <- message:
		default:
			delete(deviceSet, conn)
			close(conn.Send)
		}
	}
}

// func (h *Hub) sendGroupMessage(msg *protocol.Message) {
// 	// 发送给群组的消息，查找该群所有的用户进行发送
// 	users := h.svcCtx.GroupService.GetUserIdByGroupUuid(msg.To)
// 	for _, user := range users {
// 		if user.Uuid == msg.From {
// 			continue
// 		}

// 		client, ok := h.Clients[user.Uuid]
// 		if !ok {
// 			continue
// 		}

// 		_, err := h.svcCtx.UserModel.FindOne(client.ctx, client.UserID)
// 		// 由于发送群聊时，from是个人，to是群聊uuid。所以在返回消息时，将form修改为群聊uuid，和单聊进行统一
// 		msgSend := protocol.Message{
// 			Avatar:       "",
// 			FromUsername: msg.FromUsername,
// 			From:         msg.To,
// 			To:           msg.From,
// 			Content:      msg.Content,
// 			ContentType:  msg.ContentType,
// 			Type:         msg.Type,
// 			MessageType:  msg.MessageType,
// 			Url:          msg.Url,
// 		}

// 		msgByte, err := proto.Marshal(&msgSend)
// 		if err == nil {
// 			client.Send <- msgByte
// 		}
// 	}
// }
