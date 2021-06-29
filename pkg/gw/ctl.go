package gw

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	istio "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GateWayCtl struct {
	GwService *GateWayService  `inject:"-"`
	Client    *istio.Clientset `inject:"-"`
}

func NewGateWayCtl() *GateWayCtl {
	return &GateWayCtl{}
}

func (this *GateWayCtl) SaveGateWay(c *gin.Context) goft.Json {
	gw := &v1alpha3.Gateway{}
	goft.Error(c.ShouldBindJSON(gw))

	update := c.DefaultQuery("update", "") //前端传这个表示代表是编辑
	if update != "" {
		oldGw := this.GwService.LoadGw(gw.Namespace, gw.Name)
		gw.ResourceVersion = oldGw.ResourceVersion
		_, err := this.Client.NetworkingV1alpha3().Gateways(gw.Namespace).Update(c, gw, v1.UpdateOptions{})
		goft.Error(err)
	} else {
		_, err := this.Client.NetworkingV1alpha3().Gateways(gw.Namespace).Create(c, gw, v1.CreateOptions{})
		goft.Error(err)
	}

	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

//接收ns参数，有则显示 ns下的，没有则显示全部， 全部的话是一个map
func (this *GateWayCtl) GwList(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "")
	var ret interface{}
	if ns == "" {

		ret = this.GwService.ListAll()
	} else {
		ret = this.GwService.ListGW(ns)
	}
	return gin.H{
		"code": 20000,
		"data": ret,
	}
}

//加载网关详细
func (this *GateWayCtl) LoadGW(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")

	return gin.H{
		"code": 20000,
		"data": this.GwService.LoadGw(ns, name),
	}
}
func (this *GateWayCtl) DeleteGW(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	goft.Error(this.Client.NetworkingV1alpha3().Gateways(ns).Delete(c, name, v1.DeleteOptions{}))
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}
func (*GateWayCtl) Name() string {
	return "GateWayCtl"
}
func (this *GateWayCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/gateways", this.GwList)
	goft.Handle("POST", "/gateways", this.SaveGateWay)

	//加载网关详细
	goft.Handle("GET", "/gateways/:ns/:name", this.LoadGW)
	goft.Handle("DELETE", "/gateways/:ns/:name", this.DeleteGW)
}
