---

layout: post
title: ticker会不会堆积
date: 2018-08-09T11:35:22+0800

---

```
package main

import (
	"flag"
	"sync"
	"time"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	glog.Info("start")
	ticker := time.NewTicker(time.Second * 1)
	i := 3
	lock := &sync.Mutex{}
	for range ticker.C {
		lock.Lock()
		glog.Info(i)
		time.Sleep(time.Second * time.Duration(i))
		if i != 0 {
			i = (i + 1) % 6
		}
		lock.Unlock()
	}
}
```

```
% go run a.go --logtostderr                                                                                     1 ↵
I0809 11:34:22.535851   54904 a.go:13] start
I0809 11:34:23.537187   54904 a.go:19] 3
I0809 11:34:26.537191   54904 a.go:19] 4
I0809 11:34:30.537172   54904 a.go:19] 5
I0809 11:34:35.537402   54904 a.go:19] 0
I0809 11:34:36.539792   54904 a.go:19] 0
I0809 11:34:37.540527   54904 a.go:19] 0
I0809 11:34:38.536799   54904 a.go:19] 0
```
