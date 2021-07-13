---

date: 2021-07-13T14:49:39+0800

---

```
cat cm.json | jq '.data.results[] | select(.azone | test("^SHANGHAI")) | .ip'
```
