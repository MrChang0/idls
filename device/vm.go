package device

import (
	"github.com/yuin/gopher-lua"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
)

func pathExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return false
}

var luacode = `local idls = require("IDLS")
{{range .}}
--{{ .Description }}
function {{.Name}}{{$length := len .Args }}{{$length := dec $length}}({{range $i,$v := .Args}}{{$v.Name}}{{if ne $length $i}},{{end}}{{end}})
-- write code here
-- don't delete this function
end

{{end}}
`

func createFile(filename string, signals []SignalType) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	template.Must(template.New(path.Base(filename)).Funcs(template.FuncMap{
		"dec": func(n int) int { return n - 1 },
	}).Parse(luacode)).Execute(file, signals)
	return nil
}

func (d *Device) vminit() error {
	filename := "lua/" + d.UUID + ".lua"
	exist := pathExists(filename)
	if !exist {
		createFile(filename, d.Signals)
	}
	d.l = lua.NewState()
	d.l.PreloadModule("IDLS", Loader)
	return d.l.DoFile(filename)
}

func (d *Device) GetCode() (string, error) {
	filename := "lua/" + d.UUID + ".lua"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (d *Device) Error() string {
	return d.err
}

func (d *Device) NewCode(code string) error {
	err := d.l.DoString(code)
	if err != nil {
		return err
	}
	filename := "lua/" + d.UUID + ".lua"
	err = ioutil.WriteFile(filename, []byte(code), 0644)
	if err != nil{
		return err
	}
	return d.vminit()
}

func paramsFromType(signalType SignalType, s *Signal) []lua.LValue {
	params := make([]lua.LValue, 0, len(signalType.Args))
	for i, v := range signalType.Args {
		if s.Args[i].Data == "" {
			params = append(params, lua.LNil)
			continue
		}
		switch v.Type {
		case "string":
			{
				params = append(params, lua.LString(s.Args[i].Data))
			}
		case "number":
			{
				number, err := strconv.ParseFloat(string(s.Args[i].Data), 64)
				if err != nil {
					log.Printf("parse error param name %s", v.Name)
				}
				params = append(params, lua.LNumber(number))
			}
		case "bool":
			{
				value := strings.ToLower(s.Args[i].Data)
				switch value {
				case "true":
					params = append(params, lua.LTrue)
				case "false":
					params = append(params, lua.LFalse)
				}
			}
		}
	}
	return params
}

func (d *Device) RunSignalFunc(s *Signal) error {
	signalType, err := d.findSignalType(s.Name)
	if err != nil {
		return err
	}
	params := paramsFromType(signalType, s)
	return d.l.CallByParam(lua.P{
		Fn:      d.l.GetGlobal(s.Name),
		NRet:    0,
		Protect: true,
	}, params...)
}
