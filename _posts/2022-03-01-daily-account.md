---

date: 2022-03-01T16:42:04+0800

---

昨天中午出去打卡了一下珈琲咖啡店，吃了一份咖啡博士做的咖啡。

下午带小赛赛去医院看咳嗽，老是夜里咳嗽，白天没有。 医生也没怎么看，就说是过敏性的，开了一点抗过敏的药。昨天晚上吃了一点，晚上的确是好了不少。

今天先扩容了一个 Jenkins 的 Outband 机器，然后一直在折腾怎么把 trivy + trivy-harbor-adaptor 跑到容器里面。

```
RUN yum install -y podman
RUN sed -i 's/driver = .*/driver = "vfs"/' /etc/containers/storage.conf
RUN rm /var/lib/containers/storage/libpod/bolt_state.db || echo 'no bolt_state.db'
```
