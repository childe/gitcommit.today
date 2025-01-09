---

date: 2025-01-09T17:27:24+0800
layout: post

---

方法一：使用默认参数

```python
functions = []
for i in range(10):
    functions.append(lambda x=i: print(x))

for f in functions:
    f()
```

方法二：使用闭包

```python
def create_printer(i):
    def inner():
        print(i)
    return inner

functions = [create_printer(i) for i in range(10)]

for f in functions:
    f()
```
