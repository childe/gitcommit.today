---
layout: post
title:  "elasticsearch中限制搜索超时时间"
date:   2016-06-17 00:36:29 +0800
---

```
PUT _cluster/settings
{
   "transient": {"search.default_search_timeout":"250s"}
}
```
