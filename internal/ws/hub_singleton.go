package ws

import (
	"ginchat2/internal/svc"
	"sync"
)

var (
	hubOnce sync.Once
	hubInst *Hub
)

// GetHub returns a process-wide Hub singleton.
// Websocket connections must share the same Hub, otherwise online user routing breaks.
func GetHub(svcCtx *svc.ServiceContext) *Hub {
	hubOnce.Do(func() {
		hubInst = NewHub(svcCtx)
	})
	return hubInst
}
