---

date: 2022-06-19T23:52:14+0800
layout: post

---

```go
package main

func main() {
	arr := make([]int, 0)
	go func() {
		for i := range arr {
			print(i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			arr = append(arr, i)
		}
	}()
}
```

测试结果

```
go run -race main.go
==================
WARNING: DATA RACE
Write at 0x00c00009e000 by goroutine 7:
  main.main.func2()
      /Users/liujia/tmp/1655653776/main.go:13 +0xcb

Previous read at 0x00c00009e000 by goroutine 6:
  main.main.func1()
      /Users/liujia/tmp/1655653776/main.go:6 +0x30

Goroutine 7 (running) created at:
  main.main()
      /Users/liujia/tmp/1655653776/main.go:11 +0x14e

Goroutine 6 (finished) created at:
  main.main()
      /Users/liujia/tmp/1655653776/main.go:5 +0xe4
==================
Found 1 data race(s)
exit status 66
```
