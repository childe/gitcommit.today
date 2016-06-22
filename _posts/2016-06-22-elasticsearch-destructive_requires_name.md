---
layout: post
title:  "elasticsearch对可能破坏性的操作限制通配符"
date:   2016-06-22 14:19:03 +0800
---

```
PUT _cluster/settings?master_timeout=5m
{
   "transient": {
         "action.destructive_requires_name": true
    }
}
```
