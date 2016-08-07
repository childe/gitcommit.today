---
layout: post
title:  "golang中select是同步的"
date:   2016-08-07 12:19:18 +0800
---

```
	glog.Info(data)
	select {
	case zevent := <-watch:
		glog.Info(zevent)
	}
	glog.Info(data)
```

select 是同步的, 打印出zevent之后, 再打印后面的data
