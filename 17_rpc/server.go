package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"sync"
)

type Listener int

var num int = 0
var mt sync.Mutex

func (l *Listener) GetLine(line []byte, ack *bool) error {
	num++
	//fmt.Println(string(line))
	fmt.Println(num)
	return nil
}

func (l *Listener) AutoIncreament(tt string, ack *int) error {
	mt.Lock()
	defer mt.Unlock()
	num++
	*ack = num
	//fmt.Println(string(line))
	fmt.Println(tt, " ", num)

	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "192.168.1.23:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
