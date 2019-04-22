package main

import (
	"log"
	"strconv"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("topicwxq", "xxx", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}
	//c.MaxInFlight defaults to 1
	gCounter := 0

	c.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//log.Println("NSQ message received:")
		//log.Println(string(message.Body))
		tmpCount, _ := strconv.Atoi(string(message.Body))
		if tmpCount < gCounter {
			log.Printf("tmpCount %d is less than gCounter %d", tmpCount, gCounter)
		} else {
			gCounter = tmpCount
			log.Println(string(message.Body))
		}
		return nil
	}))

	err = c.ConnectToNSQD("192.168.1.81:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"topicwxq\"...")
	wg.Wait()
}
