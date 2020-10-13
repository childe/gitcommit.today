---

date: 2020-10-13T20:01:10+0800

---

systemtap 真好用

```
stap -e 'probe netfilter.ip.local_out {
  if (dport == 53) # or parametrize
      printf("%s[%d %d] %s:%d\n", execname(), ppid(), pid(), daddr, dport)
}'
```
