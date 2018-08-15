package device

import (
	"log"

	"github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)

	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"call": call,
}

func call(L *lua.LState) int {
	result := lua.LFalse
	defer L.Push(result)
	name := L.CheckString(1)
	top := L.GetTop()
	d, ok := CheckDeviceByName(name)
	if !ok {
		log.Println("can't find device name:", name)
		return 1
	}
	funcname := L.CheckString(2)
	argTypes, ok := d.FindEventArgs(funcname)
	if !ok {
		log.Println("can't find event from device name:", name)
		return 1
	}
	index := 3
	args := make([]Arg, 0, len(argTypes))
	for _, argType := range argTypes {
		if index > top {
			log.Printf("device:%s,event:%s,args error\n", name, funcname)
			return 1
		}
		args = append(args, Arg{argType.Name, L.Get(index).String()})
		index = index + 1
	}
	d.eventChan <- Event{funcname, args}
	result = lua.LTrue
	return 1
}