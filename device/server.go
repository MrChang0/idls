package device

import (
	"encoding/json"
	"errors"
	"github.com/MrChang0/idls/idp"
	"log"
	"net"
)

func Register(conn net.Conn) (*Device, error) {
	if err := idp.SendCommand(conn, "register", nil); err != nil {
		return nil, err
	}

	c, data, err := idp.GetCommand(conn)
	if err != nil {
		return nil, err
	}

	if c != "register" {
		return nil, errors.New("cmd error")
	}
	d := &Device{}
	err = json.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Device) SendEvent(event *Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	if err := idp.SendCommand(d.conn, "event", data); err != nil {
		return err
	}
	return nil
}
