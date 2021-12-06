---

date: 2021-12-06T15:12:09+0800

---

```
package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]sync.Mutex{}
	lock := sync.Mutex{}
	fmt.Printf("%p\n", &lock)
	m["baidu.com"] = lock

	get := m["baidu.com"]
	fmt.Printf("%p\n", &get)
}
```

```
# go run a.go
0xc0000b2008
0xc0000b2020
```
