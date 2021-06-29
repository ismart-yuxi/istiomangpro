package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//func main() {
//	client := bootstrap.NewK8sConfig().IstioRestClient()
//	list, err := client.NetworkingV1alpha3().Gateways("myistio").
//		List(context.Background(), v1.ListOptions{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, gw := range list.Items {
//		fmt.Println(gw.Name, gw.Spec.Servers[0].Hosts[0])
//	}
//}

func TestPathSeparator(t *testing.T) {
	fmt.Println(string(os.PathSeparator))
	assert.Equal(t, "\\", string(os.PathSeparator))
}
