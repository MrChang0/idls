package device

import (
	"net"

	"encoding/json"
	"errors"
	"github.com/yuin/gopher-lua"
)

type ArgType struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

type Arg struct {
	Name string `json:"name" bson:"name"`
	Data string `json:"data" bson:"data"`
}

type SignalType struct {
	Name        string    `json:"name" bson:"name" `
	Description string    `json:"description" bson:"description"`
	Args        []ArgType `json:"args" bson:"args"`
}

type Signal struct {
	Name string `json:"name" bson:"name"`
	Args []Arg  `json:"args" bson:"args"`
}

type EventType struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Args        []ArgType `json:"args" bson:"args"`
}

type Event struct {
	Name string `json:"name" bson:"name"`
	Args []Arg  `json:"args" bson:"args"`
}

type Device struct {
	UUID        string       `json:"uuid" bson:"_id"`
	Type        string       `json:"type" bson:"type"`
	Name        string       `json:"name" bson:"name"`
	Description string       `json:"description" bson:"description"`
	Signals     []SignalType `json:"signals" bson:"signals"`
	Events      []EventType  `json:"events" bson:"events"`

	conn      net.Conn `json:"-" bson:"-"`
	eventChan chan Event

	l   *lua.LState
	err string
}

func (d *Device) FindEventArgs(funcname string) ([]ArgType, bool) {
	for _, event := range d.Events {
		if event.Name == funcname {
			return event.Args, true
		}
	}
	return nil, false
}

func (d *Device) SetConn(conn net.Conn) {
	if d.conn != nil{
		d.conn.Close()
	}
	d.conn = conn
}

func (d *Device) findSignalType(name string) (SignalType, error) {
	for _, v := range d.Signals {
		if v.Name == name {
			return v, nil
		}
	}
	return SignalType{}, errors.New("can't find Signal:" + name + "from device:" + d.Name)
}

func (d *Device) EventType() (string, error) {
	b, err := json.Marshal(d.Events)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
