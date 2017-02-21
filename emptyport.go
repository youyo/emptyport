package emptyport

import (
	"net"
	"strconv"
)

func Get() (p int, err error) {
	addr, err := fetchFreeAddress()
	if err != nil {
		return 0, err
	}
	p, err = splitPort(addr)
	return
}

func fetchFreeAddress() (a string, err error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}
	defer l.Close()
	a = l.Addr().String()
	return
}

func splitPort(a string) (p int, err error) {
	_, port, err := net.SplitHostPort(a)
	if err != nil {
		return 0, err
	}
	p, err = toInt(port)
	return
}

func toInt(s string) (i int, err error) {
	i, err = strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return
}
