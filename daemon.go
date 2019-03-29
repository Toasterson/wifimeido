package wifimeido

import (
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
		_, err = c.Write(data)
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
