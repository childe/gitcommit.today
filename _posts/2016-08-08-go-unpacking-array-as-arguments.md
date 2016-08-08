---
layout: post
title:  "go-unpacking-array-as-arguments"
date:   2016-08-08 17:48:54 +0800
---

```
command := strings.Split("service nginx reload")
cmd := exec.Command(command[0], command[1:]...)
err := cmd.Run()
```
