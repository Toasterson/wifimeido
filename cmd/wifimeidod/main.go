package main

import "github.com/AlexFence/wifimeido"

func main() {
	daemon := wifimeido.NewDaemon("/tmp/wifimeido.sock")
	daemon.Run()
}
