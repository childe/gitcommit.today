---

date: 2021-06-21T14:13:22+0800

---

golang mod 里面的 replace 语法和应用简单学习一下。

## 语法

> replace module-path [module-version] => replacement-path [replacement-version]


module-version: 可以用来指定特定的版本，如果不指定，则所有版本都替换

## 例子

### 例1 换成自己的 fork

```
require example.com/othermodule v1.2.3

replace example.com/othermodule => example.com/myfork/othermodule v1.2.3-fixed
```

### 例2 换成其他版本

```
require example.com/othermodule v1.2.2

replace example.com/othermodule => example.com/othermodule v1.2.3
```

### 例3 换成本地代码

```
require example.com/othermodule v1.2.3

replace example.com/othermodule => ../othermodule
```

```
require example.com/othermodule v1.2.5

replace example.com/othermodule v1.2.5 => ../othermodule
```
