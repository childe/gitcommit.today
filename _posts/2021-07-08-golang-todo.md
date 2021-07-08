---

date: 2021-07-08T11:43:09+0800
title: 对一段代码的疑惑

---

```
conn := db.freeConn[0]
copy(db.freeConn, db.freeConn[1:])
db.freeConn = db.freeConn[:numFree-1]
```

为什么不用 freeConn = freeConn[1:]
