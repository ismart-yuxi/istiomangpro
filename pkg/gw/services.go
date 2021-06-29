package gw

import (
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

//@Service
type GateWayService struct {
	GateWayMap *GateWayMapStruct `inject:"-"`
}

func NewGateWayService() *GateWayService {
	return &GateWayService{}
}
func (this *GateWayService) ListGW(ns string) []*v1alpha3.Gateway {
	return this.GateWayMap.ListAll(ns)
}
func (this *GateWayService) ListAll() []map[string]interface{} {
	return this.GateWayMap.ListAllGateways()
}
func (this *GateWayService) LoadGw(ns, name string) *v1alpha3.Gateway {
	gw := this.GateWayMap.Get(ns, name)
	if gw == nil {
		panic("no such GateWay")
	}
	return gw
}
