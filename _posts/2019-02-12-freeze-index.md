---

title: 通过cross search访问freezed index
date: 2019-02-12T15:36:59+0800
layout: post

---

Elasticsearch 6.6.0 的文档中说, 访问 freeze 索引需要添加 `ignore_throttled=false` , 测试下来的确如此.

但是, 如果通过低版本的(6.4.1) cross search cluster 搜索, 不加这个参数居然也可以.
