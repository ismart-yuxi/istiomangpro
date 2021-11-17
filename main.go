package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"istiomang/bootstrap"
	"istiomang/pkg/gw"
	"istiomang/pkg/vs"
	"istiomang/pkg/wscore"
)

const (
	version = "/v1/"
)

//func cross() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//		if method != "" {
//			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
//			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
//			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Token")
//			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
//			c.Header("Access-Control-Allow-Credentials", "true")
//		}
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//	}
//}
func main() {
	server := goft.Ignite().Config(
		bootstrap.NewIstioHandler(),       //1
		bootstrap.NewK8sConfig(),          //2
		bootstrap.NewIstioMaps(),          //3
		bootstrap.NewIstioServiceConfig(), //4
	).Mount(version,
		vs.NewVsCtl(),
		gw.NewGateWayCtl(),
		wscore.NewWsCtl(),
	)
	server.Launch()
}
