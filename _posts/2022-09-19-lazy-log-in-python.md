---

date: 2022-09-19T12:51:47+0800
layout: post

---


```
import logging


def bad_call_with_problem():
    try:
        logging.error("haha %d" % None)
    except:
        print("Raised an error.")
        raise
    print("-" * 10)


def good_call_with_problem():
    try:
        logging.error("haha %d", None)
    except:
        print("This will never be raised.")
        raise
    print("~" * 10)


good_call_with_problem()
bad_call_with_problem()
good_call_with_problem()
```

专有名词叫 lazy logging，但我没有在官网找到这一块的说明，看到两篇文章如下

https://medium.com/swlh/why-it-matters-how-you-log-in-python-1a1085851205

https://medium.com/flowe-ita/logging-should-be-lazy-bc6ac9816906
