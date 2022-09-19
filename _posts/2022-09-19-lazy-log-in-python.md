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
