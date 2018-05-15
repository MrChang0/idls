package idp

import (
	"net"
)

func circleRead(in net.Conn, b []byte) (err error) {
	n := 0
	for {
		b = b[n:]
		length := len(b)
		n, err = in.Read(b)
		if err != nil || n == length {
			return
		}
	}
}

func sizeBuffer(b []byte) int {
	return 0xffff & (int(b[0])<<8 | int(b[1]))
}

func circleWrite(out net.Conn, b []byte) (err error) {
	n := 0
	for {
		b = b[n:]
		length := len(b)
		n, err = out.Write(b)
		if err != nil || n == length {
			return
		}
	}
}

func send(out net.Conn, b []byte) error {
	length := len(b)
	buffer := make([]byte, length+2)
	buffer[0] = byte(length>>8) & 0xff
	buffer[1] = byte(length) & 0xff
	copy(buffer[2:], b)
	return circleWrite(out, buffer)
}
