package ws

import (
	"context"
	constant "ginchat2/common"
	"ginchat2/protocol"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Client represents a websocket connection for a single authenticated user.
type Client struct {
	logger   logx.Logger
	UserID   int64
	Conn     *websocket.Conn
	Send     chan []byte
	Hub      *Hub
	cancelFn context.CancelFunc
	ctx      context.Context
}

func NewClient(ctx context.Context, w http.ResponseWriter, r *http.Request, hub *Hub, logger logx.Logger, userID int64) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		UserID:   userID,
		Conn:     conn,
		Hub:      hub,
		Send:     make(chan []byte, 256),
		logger:   logger.WithFields(logx.LogField{Key: "user", Value: userID}),
		ctx:      ctx,
		cancelFn: cancel,
	}

	hub.Register <- client
	go client.WritePump()
	go client.ReadPump()
	return client, nil
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		close(c.Send)
		c.Conn.Close()
		c.cancelFn()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			logx.Info(err)
			return
		}

		msg := &protocol.Message{}
		if err := proto.Unmarshal(message, msg); err != nil {
			c.logger.Error("Unmarshal msg fail:", err)
			continue
		}

		if msg.Type == constant.HEAT_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err := proto.Marshal(pong)
			if err != nil {
				logx.Error("client marshal message error")
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
			continue
		}

		c.Hub.Inbound <- inboundMsg{client: c, data: message}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
