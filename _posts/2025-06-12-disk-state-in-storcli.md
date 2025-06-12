---

date: 2025-06-12T11:46:41+0800
layout: post

---


总结一下，在storcli64 /c0 show all的磁盘列表中，常见的物理磁盘状态有：

```
Onln: 在线（通常指已配置在虚拟磁盘中且正常工作）
UGood: 未配置但良好
JBOD: 配置为直通磁盘
GHS: 全局热备
DHS: 专用热备
UBad: 未配置但故障
Offln: 离线（故障）
Msng: 缺失
Rbld: 重建中
Pdgd: 预测故障（即将失效）
Frn: 外来磁盘（带有其他控制器的配置）
```

