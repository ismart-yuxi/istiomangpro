package vs

import (
	"github.com/gin-gonic/gin"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istiomang/pkg/wscore"
	"log"
)

type VsHandler struct {
	VsMap     *VsMapStruct `inject:"-"`
	VsService *VsService   `inject:"-"`
}

func (this *VsHandler) OnAdd(obj interface{}) {
	this.VsMap.Add(obj.(*v1alpha3.VirtualService))
	ns := obj.(*v1alpha3.VirtualService).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "vs",
			"result": gin.H{"ns": ns,
				"data": this.VsService.ListVs(ns)},
		},
	)
}
func (this *VsHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.VsMap.Update(newObj.(*v1alpha3.VirtualService))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1alpha3.VirtualService).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "vs",
			"result": gin.H{"ns": ns,
				"data": this.VsService.ListVs(ns)},
		},
	)
}
func (this *VsHandler) OnDelete(obj interface{}) {
	this.VsMap.Delete(obj.(*v1alpha3.VirtualService))
	ns := obj.(*v1alpha3.VirtualService).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "vs",
			"result": gin.H{"ns": ns,
				"data": this.VsService.ListVs(ns)},
		},
	)
}
