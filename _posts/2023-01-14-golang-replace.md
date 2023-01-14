---

date: 2023-01-14T15:48:47+0800
layout: post

---

github.com/spf13/cast v1.4.1 有问题，想改用 oasisprotocol/cast v0.0.0-20220606122631-eba453e69641

但是直接在代码里面使用后者会报错，因为后者的代码里面还是用的 spf13/cast，和项目名字对不起来。

但是 go.mod 里面使用 replace 可以做到使用 oasisprotocol/cast

replace github.com/spf13/cast v1.4.1 => github.com/oasisprotocol/cast v0.0.0-20220606122631-eba453e69641
