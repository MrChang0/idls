package device

import (
	"sync"

	"log"
)

var devices sync.Map

var name2device = make(map[string]*Device)

func CheckDeviceByUuid(uuid string) (*Device, bool) {
	dev, ok := devices.Load(uuid)
	if ok {
		return dev.(*Device), true
	}
	return nil, false
}

func Store(d *Device) {
	if d == nil {
		return
	}
	devices.Store(d.UUID, d)
	if err := d.vminit(); err != nil {
		log.Printf(err.Error())
	}
	if d.Name != "" {
		StoreDeviceByName(d)
	}
}

func Remove(d *Device){
	if d == nil{
		return
	}
	devices.Delete(d.UUID)
	if d.Name != ""{
		delete(name2device, d.Name)
	}
}

func FindOnlineDevice(uuid string) (*Device,bool){
	d,ok := devices.Load(uuid)
	if !ok{
		return nil,false
	}
	return d.(*Device),ok
}

func GetAllOnlineDevice() []*Device {
	var ret []*Device

	devices.Range(func(key, value interface{}) bool {
		ret = append(ret, value.(*Device))
		return true
	})
	return ret
}

func StoreDeviceByName(d *Device) {
	if d.Name != ""{
		name2device[d.Name] = d
	}
}

func ChangeName(oldname string,newname string){
	d,ok := name2device[oldname]
	if !ok{
		return
	}
	delete(name2device,oldname)
	name2device[newname] = d
}

func CheckDeviceByName(name string) (d *Device, ok bool) {
	d, ok = name2device[name]
	return
}
