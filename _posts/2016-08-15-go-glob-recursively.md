---
layout: post
title:  "go get files recursively using glob?"
date:   2016-08-15 14:26:19 +0800
---

golang官方的glob不能支持**代表递归  
[https://github.com/golang/go/issues/11862](https://github.com/golang/go/issues/11862)  
[https://github.com/driskell/log-courier/issues/285](https://github.com/driskell/log-courier/issues/285)

[mattn](https://github.com/mattn)给了一个他自己写的第三方库[https://github.com/mattn/go-zglob](https://github.com/mattn/go-zglob)
