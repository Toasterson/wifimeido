package wifimeido

import (
	"encoding/json"
)

type Handler interface {
	handle() Response
}

//--------------------------------------------------------------
// Ping
//--------------------------------------------------------------
type PingHandler struct {
	cmd Command
}

func pingHandlerFromRequest(r Request) PingHandler {
	return PingHandler{cmd: r.Cmd}
}

func (p *PingHandler) handle() Response {
	result := json.RawMessage(`{"value": "pong"}`)
	return Response{Cmd: p.cmd, Result: result}
}
