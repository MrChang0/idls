package db

import (
	"log"

	"github.com/MrChang0/idls/device"
	"github.com/globalsign/mgo"
)

const url = "127.0.0.1:27017"

var (
	cdevice  *mgo.Collection
	database = "IDLS"
)

func init() {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Println("MongoDB connected error:" + err.Error())
		panic(err)
	}
	cdevice = session.DB(database).C("device")
}

func GetDevice(uuid string) (*device.Device, bool) {
	d := &device.Device{}
	err := cdevice.FindId(uuid).One(d)
	if err != nil {
		log.Println("MongoDB error:" + err.Error())
		return nil, false
	}
	return d, true
}

func RegDevice(device *device.Device) bool {
	err := cdevice.Insert(device)
	if err != nil {
		log.Println("MongoDB error:" + err.Error())
		return false
	}
	return true
}

func UpdataDevice(d *device.Device) bool{
	err := cdevice.UpdateId(d.UUID,d)
	if err!=nil {
		log.Println("MonogoDB error:"+err.Error())
		return false
	}
	return true
}