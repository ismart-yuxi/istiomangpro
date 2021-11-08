package ds

import (
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sort"
	"sync"
)

type DS []*v1alpha3.DestinationRule

func (this DS) Len() int {
	return len(this)
}
func (this DS) Less(i, j int) bool {
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this DS) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type DsMapStruct struct {
	data sync.Map // [ns string] []*Vs
}

func (this *DsMapStruct) Get(ns string, name string) *v1alpha3.DestinationRule {
	if items, ok := this.data.Load(ns); ok {
		for _, item := range items.([]*v1alpha3.DestinationRule) {

			if item.Name == name {
				return item
			}
		}
	}
	return nil
}
func (this *DsMapStruct) Add(item *v1alpha3.DestinationRule) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*v1alpha3.DestinationRule), item)
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*v1alpha3.DestinationRule{item})
	}
}
func (this *DsMapStruct) Update(item *v1alpha3.DestinationRule) error {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, range_item := range list.([]*v1alpha3.DestinationRule) {
			if range_item.Name == item.Name {
				list.([]*v1alpha3.DestinationRule)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found", item.Name)
}
func (this *DsMapStruct) Delete(svc *v1alpha3.DestinationRule) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, range_item := range list.([]*v1alpha3.DestinationRule) {
			if range_item.Name == svc.Name {
				newList := append(list.([]*v1alpha3.DestinationRule)[:i], list.([]*v1alpha3.DestinationRule)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}
func (this *DsMapStruct) ListAll(ns string) []*v1alpha3.DestinationRule {
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*v1alpha3.DestinationRule)

		sort.Sort(DS(newList)) //  按时间倒排序
		return newList
	}
	return []*v1alpha3.DestinationRule{} //返回空列表
}
