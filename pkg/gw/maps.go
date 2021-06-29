package gw

import (
	"fmt"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"sort"
	"sync"
)

type GW []*v1alpha3.Gateway

func (this GW) Len() int {
	return len(this)
}
func (this GW) Less(i, j int) bool {
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this GW) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type GateWayMapStruct struct {
	data sync.Map // [ns string] []*Gateway
}

func (this *GateWayMapStruct) Get(ns string, name string) *v1alpha3.Gateway {
	if items, ok := this.data.Load(ns); ok {
		for _, item := range items.([]*v1alpha3.Gateway) {

			if item.Name == name {
				return item
			}
		}
	}
	return nil
}
func (this *GateWayMapStruct) Add(item *v1alpha3.Gateway) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*v1alpha3.Gateway), item)
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*v1alpha3.Gateway{item})
	}
}
func (this *GateWayMapStruct) Update(item *v1alpha3.Gateway) error {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, range_item := range list.([]*v1alpha3.Gateway) {
			if range_item.Name == item.Name {
				list.([]*v1alpha3.Gateway)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found", item.Name)
}
func (this *GateWayMapStruct) Delete(svc *v1alpha3.Gateway) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, range_item := range list.([]*v1alpha3.Gateway) {
			if range_item.Name == svc.Name {
				newList := append(list.([]*v1alpha3.Gateway)[:i], list.([]*v1alpha3.Gateway)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}
func (this *GateWayMapStruct) ListAll(ns string) []*v1alpha3.Gateway {
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*v1alpha3.Gateway)

		sort.Sort(GW(newList)) //  按时间倒排序
		return newList
	}
	return []*v1alpha3.Gateway{} //返回空列表
}

//全部返回
func (this *GateWayMapStruct) ListAllGateways() []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)
	this.data.Range(func(key, value interface{}) bool {
		m := map[string]interface{}{
			"ns":   key,
			"list": value,
		}
		ret = append(ret, m)
		return true
	})

	return ret
}
