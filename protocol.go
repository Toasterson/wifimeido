package wifimeidod

import (
	"encoding/json"
)

type Command string;

const (
	PING Command = "ping"
	LIST Command = "ls"
	SETNET Command = "set-network"
)

type Request struct {
	cmd Command
	param []json.RawMessage
}

func (r *Request) validateParams() bool {
	switch r.cmd {
	case PING:
		return len(r.param) == 0
	case LIST:
		return len(r.param) == 0

	case SETNET:
		if (len(r.param) == 0) {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}
