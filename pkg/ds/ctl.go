package ds

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	"istiomang/common"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	istio "istio.io/client-go/pkg/clientset/versioned"
)

type DsCtl struct {
	DsService *DsService       `inject:"-"`
	Client    *istio.Clientset `inject:"-"`
}

func NewDsCtl() *DsCtl {
	return &DsCtl{}
}

// DsList 接收ns参数，有则显示 ns下的，没有则显示全部， 全部的话是一个map
func (this *DsCtl) DsList(c *gin.Context) goft.Json {

	ns := c.DefaultQuery("ns", "")
	return common.Success(this.DsService.ListDs(ns))
}

// SaveDS 创建和 修改
func (this *DsCtl) SaveDS(c *gin.Context) goft.Json {
	ds := &v1alpha3.DestinationRule{}
	goft.Error(c.ShouldBindJSON(ds))
	update := c.Query("update")
	if update != "" {
		old := this.DsService.LoadDs(ds.Namespace, ds.Name)
		ds.ResourceVersion = old.ResourceVersion

		_, err := this.Client.NetworkingV1alpha3().
			DestinationRules(ds.Namespace).Update(c, ds, v1.UpdateOptions{})
		goft.Error(err)
	} else {
		_, err := this.Client.NetworkingV1alpha3().
			DestinationRules(ds.Namespace).Create(c, ds, v1.CreateOptions{})
		goft.Error(err)
	}

	return common.Success(nil)
}
func (this *DsCtl) DeleteDS(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	goft.Error(this.Client.NetworkingV1alpha3().DestinationRules(ns).Delete(c, name, v1.DeleteOptions{}))
	return common.Success(nil)
}

//加载ds详细
func (this *DsCtl) LoadDS(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	return common.Success(this.DsService.LoadDs(ns, name))
}
func (*DsCtl) Name() string {
	return "DsCtl"
}

func (this *DsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/destinationrule", this.DsList)

	//保存或新增
	goft.Handle("POST", "/destinationrule", this.SaveDS)

	//加载DS详细
	goft.Handle("GET", "/destinationrule/:ns/:name", this.LoadDS)
	goft.Handle("DELETE", "/destinationrule/:ns/:name", this.DeleteDS)

}
