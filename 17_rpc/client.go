package main

import (
	"log"
	"net/rpc"
	"sync"
)

func RpcCallAutoInCreament(client *rpc.Client) {
	var reply int
	err := client.Call("Listener.AutoIncreament", "client1111", &reply)
	log.Printf("reply:%d", reply)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	client, err := rpc.Dial("tcp", "192.168.1.23:42586")
	if err != nil {
		log.Fatal(err)
	}

	var w sync.WaitGroup
	w.Add(1)
	i := 0
	//in := bufio.NewReader(os.Stdin)
	for {
		/*line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}*/

		//err = client.Call("Listener.GetLine", line, &reply)
		go RpcCallAutoInCreament(client)
		i++
		if i > 100000 {
			break
		}
	}

	w.Wait()
}
