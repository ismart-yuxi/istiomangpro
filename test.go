package main

import (
	"context"
	"fmt"
	"istiomang/bootstrap"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {
	client := bootstrap.NewK8sConfig().IstioRestClient()
	list, err := client.NetworkingV1alpha3().Gateways("myistio").
		List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, gw := range list.Items {
		fmt.Println(gw.Name, gw.Spec.Servers[0].Hosts[0])
	}

}
