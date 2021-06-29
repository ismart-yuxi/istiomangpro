package bootstrap

import (
	"istiomang/pkg/gw"
	"istiomang/pkg/vs"
)

//注入 回调handler
type IstioHandler struct{}

func NewIstioHandler() *IstioHandler {
	return &IstioHandler{}
}

// VsHandler handler
func (this *IstioHandler) VsHandler() *vs.VsHandler {
	return &vs.VsHandler{}
}

// GWHandler handler
func (this *IstioHandler) GWHandler() *gw.GateWayHandler {
	return &gw.GateWayHandler{}
}
