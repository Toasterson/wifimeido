package main

import (
	"fmt"
	"github.com/alexfence/wifimeido/networks"
)

func main() {
	net := networks.NewRegularNetwork("myNet", "123")

	storage := networks.NewFileStorage("/tmp/net/")

	err := storage.SaveNetwork(net)

	fmt.Println(err)

	var net2 networks.RegularNetwork

	storage.OpenNetwork("myNet", net2)

	fmt.Println(net2.PSK())
}
