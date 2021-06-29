package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"istiomang/common/response"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NsCtl struct {
	Client *kubernetes.Clientset `inject:"-"`
}

func NewNsCtl() *NsCtl {
	return &NsCtl{}
}

func (this *NsCtl) ListAll(c *gin.Context) {
	list, err := this.Client.CoreV1().Namespaces().List(c, v1.ListOptions{})
	goft.Error(err)

	ret := make([]*NsModel, len(list.Items))
	for index, item := range list.Items {
		istio := false
		if _, ok := item.Labels["istio-injection"]; ok {
			istio = true
		}
		ret[index] = &NsModel{Name: item.Name, Istio: istio}
	}
	response.Success(c, "命名空间", ret)
}

func (*NsCtl) Name() string {
	return "VsCtl"
}

func (this *NsCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/nslist", this.ListAll)
}