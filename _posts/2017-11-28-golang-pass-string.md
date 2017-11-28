---

layout: post
title: golang里面string是传值还是传引用
date: 2017-11-28 16:56:52 +0800

---

结论: 传值, 会copy一份新数据

```
package main

import "fmt"

func byval(q string) {
	fmt.Printf("3. byval -- q %T: &q=%p q=%v\n", q, &q, q)
	q = "xyz"
	fmt.Printf("4. byval -- q %T: &q=%p q=%v\n", q, &q, q)
}

func main() {
	i := "abcd"
	fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
	p := i
	byval(p)
	fmt.Printf("5. main  -- p %T: &p=%p p=%v\n", p, &p, p)
	fmt.Printf("6. main  -- i  %T: &i=%p i=%v\n", i, &i, i)
}
```

```
% go run a.go
1. main  -- i  string: &i=0xc42000a2e0 i=abcd
3. byval -- q string: &q=0xc42000a330 q=abcd
4. byval -- q string: &q=0xc42000a330 q=xyz
5. main  -- p string: &p=0xc42000a320 p=abcd
6. main  -- i  string: &i=0xc42000a2e0 i=abcd
```
