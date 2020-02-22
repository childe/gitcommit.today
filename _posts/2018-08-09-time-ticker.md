---

layout: post
title: ticker会不会堆积
date: 2018-08-09T11:35:22+0800

---

```
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
	ticker := time.NewTicker(time.Second * 3)
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
```

```
2020/02/22 19:14:03 start
2020/02/22 19:14:06 sleep 3 second
2020/02/22 19:14:09 sleep 4 second
2020/02/22 19:14:13 sleep 5 second
2020/02/22 19:14:18 sleep 0 second
2020/02/22 19:14:21 sleep 1 second
2020/02/22 19:14:24 sleep 2 second
2020/02/22 19:14:27 sleep 3 second
2020/02/22 19:14:30 sleep 4 second
2020/02/22 19:14:34 sleep 5 second
```
