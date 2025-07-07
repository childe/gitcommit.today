package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

func main() {
	flag.Parse()
	log.Println("start")
	ticker := time.NewTicker(time.Second * 2)
	i := 3
	lock := &sync.Mutex{}
	for range ticker.C {
		lock.Lock()
		log.Printf("sleep %d second", i)
		time.Sleep(time.Second * time.Duration(i))
		i = (i + 1) % 6
		lock.Unlock()
	}
}

