---

date: 2022-07-07T15:36:52+0800
layout: post

---

[https://programmer.group/kafka-golang-client-introduction.html](https://programmer.group/kafka-golang-client-introduction.html)

看到一篇文章，里面有一点点提到了 healer，这介绍还真的是非常准确。

```
Client name	Advantages and disadvantages
sarama	The number of users is relatively large, but it is relatively difficult to use, with better performance
confluent-kafka-go	The encapsulation of Kafka in C language has strong performance, but depends on librdkafka
kafka-go	The operation is simple, but the performance is poor. The common dual core CPU machine in the production environment only processes about 300 pieces per second
healer	The operation is very simple, and the performance is similar to that of sarama. This product is the work of a big bull of Ctrip. At present, the number of users in the community is small, and there is no corresponding support and maintenance
```
