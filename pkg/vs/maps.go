package vs

import (
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sort"
	"sync"
)

type VS []*v1alpha3.VirtualService

func (this VS) Len() int {
	return len(this)
}
func (this VS) Less(i, j int) bool {
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this VS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type VsMapStruct struct {
	data sync.Map // [ns string] []*Vs
}

func (this *VsMapStruct) Get(ns string, name string) *v1alpha3.VirtualService {
	if items, ok := this.data.Load(ns); ok {
		for _, item := range items.([]*v1alpha3.VirtualService) {

			if item.Name == name {
				return item
			}
		}
	}
	return nil
}
func (this *VsMapStruct) Add(item *v1alpha3.VirtualService) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*v1alpha3.VirtualService), item)
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*v1alpha3.VirtualService{item})
	}
}
func (this *VsMapStruct) Update(item *v1alpha3.VirtualService) error {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, range_item := range list.([]*v1alpha3.VirtualService) {
			if range_item.Name == item.Name {
				list.([]*v1alpha3.VirtualService)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found", item.Name)
}
func (this *VsMapStruct) Delete(svc *v1alpha3.VirtualService) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, range_item := range list.([]*v1alpha3.VirtualService) {
			if range_item.Name == svc.Name {
				newList := append(list.([]*v1alpha3.VirtualService)[:i], list.([]*v1alpha3.VirtualService)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}
func (this *VsMapStruct) ListAll(ns string) []*v1alpha3.VirtualService {
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*v1alpha3.VirtualService)

		sort.Sort(VS(newList)) //  按时间倒排序
		return newList
	}
	return []*v1alpha3.VirtualService{} //返回空列表
}
