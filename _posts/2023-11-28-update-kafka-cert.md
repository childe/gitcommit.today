---

date: 2023-11-28T14:12:38+0800
layout: post

---

在重启 ZK 之前先更新 broker 里面的证书并 reload

```
export c="" && for i in {0..9} ; do k exec -ti $c-shaxy-b-kafka-$i  -npro-kafka -- sh -c 'export CERTS_STORE_PASSWORD=$(grep listener.name.replication-9091.ssl.keystore.password /tmp/strimzi.properties | cut -f2 -d=) && sh /opt/kafka/kafka_tls_prepare_certificates.sh' ; k exec -ti $c-shaxy-b-kafka-$i -npro-kafka -- bin/kafka-configs.sh --bootstrap-server 127.0.0.1:9092  --entity-type brokers --entity-name $i --alter --add-config listener.name.replication-9091.ssl.truststore.location=/tmp/kafka/cluster.truststore.p12 ; done
```

这个命令应该是会去重载证书（虽然路径名字未变），因为中间尝试删除 truststore 再运行命令行报错，说文件不存在。

而且还做了一个“破坏性”的测试：不运行 kafka-config.sh 直接重启一个 Broker，会导致其他 Broker 全部报证书验证失败。然后再刷一下 kafka-confg.sh，报错停止。所以看起来这个脚本的确是重载了新的证书。

按这个方法跑下来，RB 全部无损重启了。

但 XY 有大概6个集群还是出问题，需要重启 Broker 才行。

这些需要重启的 Broker 在支行上面的 kafka-config 命令时报错如下：

```
Error while executing config command with args '--bootstrap-server 127.0.0.1:9092 --entity-type brokers --entity-name 0 --alter --add-config listener.name.replication-9091.ssl.truststore.location=/tmp/kafka/cluster.truststore.p12'
java.util.concurrent.ExecutionException: org.apache.kafka.common.errors.InvalidRequestException: Invalid config value for resource ConfigResource(type=BROKER, name='0'): Validation of dynamic config update of SSLFactory failed: javax.net.ssl.SSLHandshakeException: PKIX path validation failed: java.security.cert.CertPathValidatorException: validity check failed
	at org.apache.kafka.common.internals.KafkaFutureImpl.wrapAndThrow(KafkaFutureImpl.java:45)
	at org.apache.kafka.common.internals.KafkaFutureImpl.access$000(KafkaFutureImpl.java:32)
	at org.apache.kafka.common.internals.KafkaFutureImpl$SingleWaiter.await(KafkaFutureImpl.java:104)
	at org.apache.kafka.common.internals.KafkaFutureImpl.get(KafkaFutureImpl.java:272)
	at kafka.admin.ConfigCommand$.alterConfig(ConfigCommand.scala:345)
	at kafka.admin.ConfigCommand$.processCommand(ConfigCommand.scala:297)
	at kafka.admin.ConfigCommand$.main(ConfigCommand.scala:90)
	at kafka.admin.ConfigCommand.main(ConfigCommand.scala)
Caused by: org.apache.kafka.common.errors.InvalidRequestException: Invalid config value for resource ConfigResource(type=BROKER, name='0'): Validation of dynamic config update of SSLFactory failed: javax.net.ssl.SSLHandshakeException: PKIX path validation failed: java.security.cert.CertPathValidatorException: validity check failed
```
