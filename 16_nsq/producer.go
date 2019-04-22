package main

import (
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
)

func produce(p *nsq.Producer, topic string, wg *sync.WaitGroup) {
	i := 0
	st := time.Now()
	for {
		/*if 1 == t {
			time.Sleep(time.Second)
		} else {
			time.Sleep(2 * time.Second)
		}*/
		if i > 10000 {
			break
		}

		time.Sleep(time.Millisecond)

		err := p.Publish(topic, []byte(strconv.Itoa(i)))
		if err != nil {
			log.Panic(err)
		}
		i++
	}

	elapsed := time.Since(st)
	log.Printf("consume time: %s ", elapsed)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	config := nsq.NewConfig()
	p, err := nsq.NewProducer("192.168.1.81:4150", config)
	if err != nil {
		log.Panic(err)
	}

	wg.Add(1)
	go produce(p, "topicwxq", wg)

	wg.Add(1)
	go produce(p, "topictest", wg)

	wg.Wait()
}
