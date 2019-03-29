package main

import "github.com/alexfence/wifimeido"

func main() {
	daemon := wifimeido.NewDaemon("/tmp/wifimeido.sock")
	daemon.Run()
}
