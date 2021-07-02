---

date: 2021-07-03T00:20:33+0800

---

```
echo | openssl s_client -servername www.example.com -connect www.example.com:443 2>/dev/null | openssl x509 -text
```
