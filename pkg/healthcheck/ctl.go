package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"istiomang/common"
)

type HealthCheckCtl struct {
}

func NewHealthCheckCtl() *HealthCheckCtl {
	return &HealthCheckCtl{}
}

func (this *HealthCheckCtl) HealthCheckHandler(c *gin.Context) goft.Json {
	return common.Success("health checked!")
}

func (*HealthCheckCtl) Name() string {
	return "HealthCheckCtl"
}

func (this *HealthCheckCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/healthCheckHandler", this.HealthCheckHandler)
}
