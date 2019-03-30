package main

import (
	"fmt"
	"github.com/alexfence/wifimeido/networks"
)

func main() {
	net := networks.NewRegularNetwork("myNet", "123")

	storage := networks.KeyringStorage{}

	err := storage.SaveNetwork(net)

	fmt.Println(err)
}
