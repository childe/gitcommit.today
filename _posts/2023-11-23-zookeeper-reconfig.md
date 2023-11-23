---

date: 2023-11-23T10:59:52+0800
layout: post

---

语法： `server.<positive id> = <address1>:<port1>:<port2>[:role];[<client port address>:]<client port>`

[https://zookeeper.apache.org/doc/r3.5.2-alpha/zookeeperReconfig.html](https://zookeeper.apache.org/doc/r3.5.2-alpha/zookeeperReconfig.html)

> The client port specification is to the right of the semicolon. The client port address is optional, and if not specified it defaults to "0.0.0.0". As usual, role is also optional, it can be participant or observer (participant by default).

前提：

1. ZK 启动的时候需要有一个 admin 用户，Environment="SERVER_JVMFLAGS=-Dzookeeper.DigestAuthenticationProvider.superDigest=admin:cGFzc3dvcmQK" 通过启动的时候添加这个参数

2. `addauth digest admin:password` zk shell 里面配置认证，然后再 reconfig
