---

date: 2024-02-19T19:18:06+0800
layout: post

---


```bash
export idx=3
export cluster=hotel-searchservice-shaxy-b

k exec -npro-kafka $cluster-kafka-$idx -- mkdir /tmp/broker-certs /tmp/client-ca-certs /tmp/cluster-ca-certs

clustercert=`k get secret -npro-kafka $cluster-cluster-ca-cert -o jsonpath='{.data.ca\.crt}' | base64 -d`
k exec -npro-kafka $cluster-kafka-$idx -- sh -c "echo '$clustercert' > /tmp/cluster-ca-certs/ca.crt"

clientscert=`k get secret -npro-kafka $cluster-clients-ca-cert -o jsonpath='{.data.ca\.crt}' | base64 -d`
k exec -npro-kafka $cluster-kafka-$idx -- sh -c "echo '$clientscert' > /tmp/client-ca-certs/ca.crt"

brokerkey=`k get secret -npro-kafka $cluster-kafka-brokers -o jsonpath="{.data.$cluster-kafka-$idx\.key}" | base64 -d`
brokercert=`k get secret -npro-kafka $cluster-kafka-brokers -o jsonpath="{.data.$cluster-kafka-$idx\.crt}" | base64 -d`
k exec -npro-kafka $cluster-kafka-$idx -- sh -c "echo '$brokerkey' > /tmp/broker-certs/$cluster-kafka-$idx.key"
k exec -npro-kafka $cluster-kafka-$idx -- sh -c "echo '$brokercert' > /tmp/broker-certs/$cluster-kafka-$idx.crt"


k exec -npro-kafka $cluster-kafka-$idx -- sh -c "sed 's|/opt/kafka/|/tmp/|' kafka_tls_prepare_certificates.sh > /tmp/upate-cert.sh"

k exec -npro-kafka $cluster-kafka-$idx -- sh -c 'export CERTS_STORE_PASSWORD=$(grep listener.name.replication-9091.ssl.keystore.password /tmp/strimzi.properties | cut -f2 -d=) && sh /tmp/upate-cert.sh'
k exec -npro-kafka $cluster-kafka-$idx -- bin/kafka-configs.sh --bootstrap-server 127.0.0.1:9092  --entity-type brokers --entity-name $idx --alter --add-config listener.name.replication-9091.ssl.truststore.location=/tmp/kafka/cluster.truststore.p12
```
