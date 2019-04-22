package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrementMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incrementChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	//ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go incrementMutex(&w, &m)
		//go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
