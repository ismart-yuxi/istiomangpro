package wscore

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type WsClient struct {
	conn      *websocket.Conn
	readChan  chan *WsMessage //读队列 (chan)
	closeChan chan byte       // 失败队列
	Locker    sync.Mutex      //加锁 。目前ws不支持并发写
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn, readChan: make(chan *WsMessage), closeChan: make(chan byte)}
}
func (this *WsClient) Ping(waittime time.Duration) {
	for {
		time.Sleep(waittime)
		func() {
			this.Locker.Lock()
			defer this.Locker.Unlock()
			err := this.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
			if err != nil {
				ClientMap.Remove(this.conn)
				return
			}
		}()

	}
}
func (this *WsClient) ReadLoop() {
	for {
		t, data, err := this.conn.ReadMessage()
		if err != nil {
			_ = this.conn.Close()
			ClientMap.Remove(this.conn)
			this.closeChan <- 1
			break
		}
		this.readChan <- NewWsMessage(t, data)
	}
}
