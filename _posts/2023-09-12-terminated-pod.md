---

date: 2023-09-12T10:42:55+0800
layout: post

---

delete pod 卡在 “I0911 18:34:01.334912  825565 kubelet_pods.go:979] Pod "xxxx(689dec0e-6057-4221-aa29-d9618c20d3c7)" is terminated, but some volumes have not been cleaned up”。

mount | grep 689dec0e

看到有一些相关的 mount 需要 umount。
