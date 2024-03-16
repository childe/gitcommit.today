---

date: 2024-03-16T10:31:35+0800
layout: post

---

```go
package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func processIsTranslated() (bool, error) {
	// https://developer.apple.com/documentation/apple_silicon/about_the_rosetta_translation_environment#3616845
	ret, err := unix.SysctlUint32("sysctl.proc_translated")
	print(ret)
	print(err)

	if err == nil {
		return ret == 1, nil
	} else if err.(unix.Errno) == unix.ENOENT {
		return false, nil
	}

	return false, err
}

func main() {
	if isTranslated, err := processIsTranslated(); err != nil {
		panic(err)
	} else if isTranslated {
		fmt.Println("Running on Rosetta 2")
	} else {
		fmt.Println("Running natively")
	}
}
```
