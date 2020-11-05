---

date: 2020-11-05T15:53:06+0800

---

systemctl start docker 的时候, 总是 restoring container , 但又恢复不了, 一直卡着.

/var/run/containerd/io.containerd.runtime.v1.linux/ 在这里找一下 container 目录, 删掉重启就好了
