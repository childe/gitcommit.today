---

date: 2022-03-04T10:59:20+0800

---

harbor 里面 Robot 帐户不能用来做 Replication。可以下面这样 WorkAround。

```
insert into harbor_user (username,password,realname) values ('robot$replication','','robot$replication');
```
