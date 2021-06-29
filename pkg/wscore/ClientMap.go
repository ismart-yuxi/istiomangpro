package wscore

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

//外部公共使用,保存所有客户端对象的Map
var ClientMap *ClientMapStruct

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map //  key 是客户端IP  value 就是 WsClient连接对象
	lock sync.Mutex
}

func (this *ClientMapStruct) Store(conn *websocket.Conn) {
	wsClient := NewWsClient(conn)
	this.data.Store(conn.RemoteAddr().String(), wsClient)
	go wsClient.Ping(time.Second * 1)
	go wsClient.ReadLoop() //处理读 循环
	// go wsClient.HandlerLoop() //处理 总控制循环
}

//向所有客户端 发送消息--发送资源列表
func (this *ClientMapStruct) SendAll(v interface{}) {
	this.data.Range(func(key, value interface{}) bool {
		//this.lock.Lock() //这里加个锁，因为目前不支持并发写wsClient
		//defer this.lock.Unlock()
		func() {
			c := value.(*WsClient).conn
			value.(*WsClient).Locker.Lock()
			defer value.(*WsClient).Locker.Unlock()
			err := c.WriteJSON(v)
			if err != nil {
				this.Remove(c)
				log.Println(err)
			}
		}()

		return true
	})
}
func (this *ClientMapStruct) Remove(conn *websocket.Conn) {
	this.data.Delete(conn.RemoteAddr().String())
}
