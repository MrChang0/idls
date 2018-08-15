package main

import (
	"log"
	"net"

	"github.com/MrChang0/idls/db"
	"github.com/MrChang0/idls/device"
	"github.com/MrChang0/idls/web"
)

func GetDevice(conn net.Conn) (d *device.Device, err error) {
	uuid, err := device.Online(conn)
	if err != nil {
		return nil, err
	}
	log.Println("uuid: " + uuid + " online")
	d, ok := device.CheckDeviceByUuid(uuid)
	if ok {
		d.SetConn(conn)
		return d, nil
	}
	d, ok = db.GetDevice(uuid)
	if ok {
		device.Store(d)
		d.SetConn(conn)
		return d, nil
	}
	d, err = device.Register(conn)
	if err != nil {
		return nil, err
	}

	if ok = db.RegDevice(d); !ok {
		log.Println("DB reg fail")
	}
	device.Store(d)
	d.SetConn(conn)
	return d, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:5566")
	if err != nil {
		log.Fatalf("listen error")
	}
	go func(l net.Listener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err.Error())
			}
			go func(conn net.Conn) {
				d, err := GetDevice(conn)
				if err != nil {
					log.Printf("device not found:%s", err.Error())
					conn.Close()
					return
				}
				d.Start()
			}(conn)
		}
	}(listen)
	web.Run()
}
