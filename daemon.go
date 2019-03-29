package wifimeido

import (
	"encoding/json"
	"log"
	"net"
)

type Daemon struct {
	socket string
}

func NewDaemon(socket string) Daemon {
	return Daemon{socket}
}

func (self Daemon) echo(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]

		println("Server got:", string(data))

		var req Request

		if err := json.Unmarshal(data, &req); err != nil {
			panic(err)
		}

		response := req.handle()

		resp_json, _ := json.Marshal(response)

		log.Print(response)

		_, err = c.Write(resp_json)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func (self *Daemon) Run() {
	l, err := net.Listen("unix", self.socket)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go self.echo(fd)
	}
}
