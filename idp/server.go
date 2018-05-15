package idp

import (
	"encoding/json"
	"net"
)

type Server struct {
	Proto   string          `json:"proto"`
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

func SendCommand(conn net.Conn, cmd string, data []byte) error {
	server := &Server{Proto: "idp.v1", Command: cmd, Data: data}
	buffer, err := json.Marshal(server)
	if err != nil {
		return err
	}
	return send(conn, buffer)
}
