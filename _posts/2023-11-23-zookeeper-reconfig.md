---

date: 2023-11-23T10:59:52+0800
layout: post

---

语法： `server.<positive id> = <address1>:<port1>:<port2>[:role];[<client port address>:]<client port>`

[https://zookeeper.apache.org/doc/r3.5.2-alpha/zookeeperReconfig.html](https://zookeeper.apache.org/doc/r3.5.2-alpha/zookeeperReconfig.html)

> The client port specification is to the right of the semicolon. The client port address is optional, and if not specified it defaults to "0.0.0.0". As usual, role is also optional, it can be participant or observer (participant by default).

client port 一般就是 2181

但这个 /zookeeper/config 路径看起来默认是有权限控制的，直接 reconfig -add 会因为没权限而失败。

有两个办法

## 方法1

1. ZK 启动的时候需要有一个 admin 用户，Environment="SERVER_JVMFLAGS=-Dzookeeper.DigestAuthenticationProvider.superDigest=admin:cGFzc3dvcmQK" 通过启动的时候添加这个参数

2. `addauth digest admin:password` zk shell 里面配置认证，然后再 reconfig

## 方法2

所有ZK 配置文件里面添加 `skipACL=yes` 然后依次重启。


另外想吐槽一点点，觉得有些和普遍的观念相悖的地方：

添加一个新 ZK 节点的时候，需要在 zoo.cfg 里面先配置上 server.xxx=xxx:xxx，否则启动失败。

然而，添加这个配置启动之后，dynamic 配置又自动把它去掉了，还需要在 zkshell 里面再次 reconfig -add 添加一次。
