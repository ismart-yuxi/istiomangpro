package bootstrap

import (
	"istiomang/pkg/gw"
	"istiomang/pkg/vs"
)

type IstioMaps struct {
}

func NewIstioMaps() *IstioMaps {
	return &IstioMaps{}
}

//初始化 VsMapStruct
func (this *IstioMaps) InitVsMap() *vs.VsMapStruct {
	return &vs.VsMapStruct{}
}

//初始化 GwMapStruct
func (this *IstioMaps) InitGwMap() *gw.GateWayMapStruct {
	return &gw.GateWayMapStruct{}
}
