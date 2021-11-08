package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
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

func TestCommander(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out: %s\n err: %s\n", outStr, errStr)
}
