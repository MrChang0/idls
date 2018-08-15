package idp

import (
"encoding/json"
"log"
"net"
)

type Client struct {
	Proto     string          `json:"proto"`
	Prototype string          `json:"prototype"`
	Data      json.RawMessage `json:"data"`
}

func getclient(conn net.Conn) (*Client, error) {
	buffer := make([]byte, 2)
	err := circleRead(conn, buffer)
	if err != nil {
		log.Printf("net read err:%s", err.Error())
	}
	length := sizeBuffer(buffer)
	buffer = make([]byte, length)
	err = circleRead(conn, buffer)
	var c Client
	if err = json.Unmarshal(buffer, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

func GetCommand(conn net.Conn) (string, []byte, error) {
	c, err := getclient(conn)
	if err != nil {
		return "", nil, err
	}
	return c.Prototype, c.Data, nil
}
