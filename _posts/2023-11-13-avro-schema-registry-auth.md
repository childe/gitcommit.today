---

date: 2023-11-13T17:54:53+0800
title: Avro Schema Registry认证头配置
layout: post

---

```sh
cat schema-registry.properties
authority.header=schema-registry-admin

curl -H 'authority.header:schema-registry.properties' -XDELETE
```

