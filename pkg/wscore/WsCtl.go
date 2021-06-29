package wscore

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"log"
)

//@Controller
type WsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
	Config *rest.Config          `inject:"-"`
}

func NewWsCtl() *WsCtl {
	return &WsCtl{}
}

func (this *WsCtl) Connect(c *gin.Context) (v goft.Void) {
	client, err := Upgrader.Upgrade(c.Writer, c.Request, nil) //升级
	if err != nil {
		log.Println(err)
		return
	} else {
		ClientMap.Store(client)
		return
	}
}

func (this *WsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/ws", this.Connect)
}
func (this *WsCtl) Name() string {
	return "WsCtl"
}
