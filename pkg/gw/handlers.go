package gw

import (
	"github.com/gin-gonic/gin"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istiomang/pkg/wscore"
	"log"
)

type GateWayHandler struct {
	GateWayMap     *GateWayMapStruct `inject:"-"`
	GateWayService *GateWayService   `inject:"-"`
}

func (this *GateWayHandler) OnAdd(obj interface{}) {
	this.GateWayMap.Add(obj.(*v1alpha3.Gateway))
	ns := obj.(*v1alpha3.Gateway).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "gw",
			"result": gin.H{"ns": ns,
				"data": this.GateWayService.ListGW(ns)},
		},
	)
}
func (this *GateWayHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.GateWayMap.Update(newObj.(*v1alpha3.Gateway))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1alpha3.Gateway).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "gw",
			"result": gin.H{"ns": ns,
				"data": this.GateWayService.ListGW(ns)},
		},
	)
}
func (this *GateWayHandler) OnDelete(obj interface{}) {
	this.GateWayMap.Delete(obj.(*v1alpha3.Gateway))
	ns := obj.(*v1alpha3.Gateway).Namespace

	wscore.ClientMap.SendAll(
		gin.H{
			"type": "gw",
			"result": gin.H{"ns": ns,
				"data": this.GateWayService.ListGW(ns)},
		},
	)
}
