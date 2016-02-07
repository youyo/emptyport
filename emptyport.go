package emptyport

import (
	"log"
	"net"
	"strconv"
)

func Get() int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
		panic("Can't listen empty port.")
	}
	addr := l.Addr().String()
	_, port, err := net.SplitHostPort(addr)
	p, _ := strconv.Atoi(port)
	defer l.Close()
	return p
}
