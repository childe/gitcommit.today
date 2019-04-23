---

title: MAC上面绑定一个新的IP到网卡
date: 2019-04-23T13:19:02+0800
layout: post

---

[https://stackoverflow.com/questions/87442/virtual-network-interface-in-mac-os-x](https://stackoverflow.com/questions/87442/virtual-network-interface-in-mac-os-x)

```
The loopback adapter is always up.

ifconfig lo0 alias 172.16.123.1 will add an alias IP 172.16.123.1 to the loopback adapter

ifconfig lo0 -alias 172.16.123.1 will remove it
```
