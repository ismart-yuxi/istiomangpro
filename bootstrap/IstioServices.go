package bootstrap

import (
	"istiomang/pkg/ds"
	"istiomang/pkg/gw"
	"istiomang/pkg/vs"
)

//@Config
type IstioServiceConfig struct{}

func NewIstioServiceConfig() *IstioServiceConfig {
	return &IstioServiceConfig{}
}
func (*IstioServiceConfig) VsService() *vs.VsService {
	return vs.NewVsService()
}
func (*IstioServiceConfig) GwService() *gw.GateWayService {
	return gw.NewGateWayService()
}

func (*IstioServiceConfig) DsService() *ds.DsService {
	return ds.NewDsService()
}
