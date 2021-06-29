package vs

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	istio "istio.io/client-go/pkg/clientset/versioned"
	"istiomang/common/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VsCtl struct {
	VsService *VsService       `inject:"-"`
	Client    *istio.Clientset `inject:"-"`
}

func NewVsCtl() *VsCtl {
	return &VsCtl{}
}

func (this *VsCtl) VsList(c *gin.Context) {
	ns := c.DefaultQuery("ns", "default")
	response.Success(c, "查询虚拟服务", this.VsService.ListVs(ns))
}

func (this *VsCtl) DeleteVS(c *gin.Context) {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "name")
	err := this.Client.NetworkingV1alpha3().VirtualServices(ns).Delete(c, name, v1.DeleteOptions{})
	goft.Error(err)
	response.Success(c, "删除虚拟服务", this.VsService.ListVs(ns))
}

//
func (this *VsCtl) SaveVS(c *gin.Context) {
	//同时处理 创建或更新操作
	isupdate := c.DefaultQuery("update", "")
	vs := &v1alpha3.VirtualService{}
	goft.Error(c.ShouldBindJSON(vs))
	if isupdate == "" { //新增
		_, err := this.Client.NetworkingV1alpha3().VirtualServices(vs.Namespace).Create(c, vs, v1.CreateOptions{})
		goft.Error(err)
	} else {
		//更新
		//要先获取原有对象
		oldVS := this.VsService.LoadVs(vs.Namespace, vs.Name) //原来的对象
		vs.ResourceVersion = oldVS.ResourceVersion            //资源版本
		_, err := this.Client.NetworkingV1alpha3().VirtualServices(vs.Namespace).Update(c, vs, v1.UpdateOptions{})
		goft.Error(err)
	}
	response.Success(c, "保存虚拟服务", nil)
}

func (this *VsCtl) VsDetail(c *gin.Context) {
	ns := c.Param("ns")
	name := c.Param("name")
	response.Success(c, "虚拟服务详情", this.VsService.LoadVs(ns, name))
}

func (*VsCtl) Name() string {
	return "VsCtl"
}

func (this *VsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/virtualservices", this.VsList)
	goft.Handle("POST", "/virtualservices", this.SaveVS)
	goft.Handle("DELETE", "/virtualservices", this.DeleteVS)

	//虚拟服务详细
	goft.Handle("GET", "/virtualservices/:ns/:name", this.VsDetail)

}
