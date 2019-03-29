package wifimeido

import (
	"encoding/json"
)

type Command string

const (
	PING   Command = "ping"
	LIST   Command = "ls"
	SETNET Command = "set-network"
)

type Request struct {
	Cmd   Command           `json:"cmd"`
	Param []json.RawMessage `json:"param"`
}

type Response struct {
	Cmd    Command         `json:"cmd"`
	Result json.RawMessage `json:"result"`
}

func (r *Request) validateParams() bool {
	switch r.Cmd {
	case PING:
		return len(r.Param) == 0
	case LIST:
		return len(r.Param) == 0

	case SETNET:
		if len(r.Param) == 0 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func (self Request) handle() Response {
	if self.validateParams() {
		switch self.Cmd {
		case PING:
			handler := pingHandlerFromRequest(self)
			return handler.handle()
		default:
			return Response{self.Cmd, nil}
		}
	} else {
		return Response{self.Cmd, nil}
	}
}
