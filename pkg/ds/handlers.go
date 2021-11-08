package ds

import (
	"github.com/gin-gonic/gin"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istiomang/pkg/wscore"
	"log"
)

type DsHandler struct {
	DsMap     *DsMapStruct `inject:"-"`
	DsService *DsService   `inject:"-"`
}

func (this *DsHandler) OnAdd(obj interface{}) {
	this.DsMap.Add(obj.(*v1alpha3.DestinationRule))
	ns := obj.(*v1alpha3.DestinationRule).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "ds",
			"result": gin.H{"ns": ns,
				"data": this.DsMap.ListAll(ns)},
		},
	)
}
func (this *DsHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.DsMap.Update(newObj.(*v1alpha3.DestinationRule))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1alpha3.DestinationRule).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "ds",
			"result": gin.H{"ns": ns,
				"data": this.DsService.ListDs(ns)},
		},
	)
}
func (this *DsHandler) OnDelete(obj interface{}) {
	this.DsMap.Delete(obj.(*v1alpha3.DestinationRule))
	ns := obj.(*v1alpha3.DestinationRule).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "ds",
			"result": gin.H{"ns": ns,
				"data": this.DsMap.ListAll(ns)},
		},
	)
}
