---

date: 2021-06-18T17:27:54+0800

---

glog 也没有使用 append 模式打开日志文件，导致第三方 truncate 日志文件之后，glog 还继续原来的 offset 写。
