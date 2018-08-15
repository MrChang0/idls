package device

import (
	"encoding/json"
	"github.com/MrChang0/idls/idp"
	"log"
)

func signal(d *Device, data []byte) {
	s := &Signal{}
	json.Unmarshal(data, s)
	log.Println("uuid:"+d.UUID+" run signal:", s.Name)
	if err := d.RunSignalFunc(s); err != nil {
		d.err = err.Error()
	}
}

func offline(d *Device, data []byte) {
	Remove(d)
}

func online(d *Device, data []byte) {
	Store(d)
}

var cmds = map[string]func(*Device, []byte){
	"signal":  signal,
	"online":  online,
	"offline": offline,
}

func (d *Device) Start() {
	d.eventChan = make(chan Event)
	go func(d *Device) {
		for {
			event, ok := <-d.eventChan
			log.Printf("%s:event name:%s",d.Name,event.Name)
			if !ok {
				log.Println("channel close")
				break
			}
			d.SendEvent(&event)
		}
	}(d)
	go func(d *Device) {
		for {
			c, data, err := idp.GetCommand(d.conn)
			log.Println("uuid:"+d.UUID+" read command "+c)
			if err != nil {
				log.Println(d.Name + " read err " + err.Error())
				offline(d, nil)
				break
			}
			go cmds[c](d, data)
		}
	}(d)

}
