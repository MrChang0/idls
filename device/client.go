package device

import (
	"encoding/json"
	"errors"
	"github.com/MrChang0/idls/idp"
	"net"
)

func Online(conn net.Conn) (string, error) {
	cmd, data, err := idp.GetCommand(conn)
	if err != nil {
		return "", err
	}

	if cmd != "online" {
		return "", errors.New("prototype need online first")
	}
	type onlinecmd struct {
		Uuid string `json:"uuid"`
	}
	o := onlinecmd{}
	err = json.Unmarshal(data, &o)
	if err != nil {
		return "", err
	}
	return o.Uuid, nil
}
