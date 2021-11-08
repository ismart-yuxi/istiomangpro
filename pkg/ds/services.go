package ds

import (
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
)

//@Service
type DsService struct {
	DsMap *DsMapStruct `inject:"-"`
}

func NewDsService() *DsService {
	return &DsService{}
}
func (this *DsService) ListDs(ns string) []*v1alpha3.DestinationRule {
	return this.DsMap.ListAll(ns)
}
func (this *DsService) LoadDs(ns, name string) *v1alpha3.DestinationRule {
	ds := this.DsMap.Get(ns, name)
	if ds == nil {
		panic("no such DestinationRule")
	}
	return ds
}
