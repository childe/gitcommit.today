---

date: 2023-11-23T17:39:42+0800
layout: post

---

刚刚发现一个现象，正常吗？

如果一个Topic是下面这样的，它不会 unclean leader elect 选出来0 做leader，需要再次触发一下 unclean.leader.election.enable=true，哪怕这个 topic 已经是 true

```json
{
  "PartitionErrorCode": 72,
  "PartitionID": 43,
  "Leader": -1,
  "LeaderEpoch": 21,
  "Replicas": [
    0,
    3
  ],
  "Isr": [
    3
  ],
  "OfflineReplicas": [
    3
  ]
}
```
