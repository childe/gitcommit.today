---

date: 2018-12-11T17:57:14+0800
layout: post
title: elastic reindex

---

```
POST _reindex?
{
  "source": {
    "index": "kibana-int"
  },
  "dest": {
    "index": "kibana-int-6.4",
    "type": "dashboard"
  },
  "script": {
    "source": "if (ctx._source.type == 'temp') {ctx._source['type']='temp'} else {ctx._source['type']='dashboard'}",
    "lang": "painless"
  }
}
```
