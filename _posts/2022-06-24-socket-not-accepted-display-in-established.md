---

date: 2022-06-24T13:15:09+0800
title: 没有Accept的Socket是不是显示在ESTABLISHED
layout: post

---

在 MACOS 测试，会显示。


```python
import socket
import time

HOST = ""  # Symbolic name meaning all available interfaces
PORT = 50007  # Arbitrary non-privileged port
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind((HOST, PORT))
    s.listen(1)
    time.sleep(10)
    conn, addr = s.accept()
    with conn:
        print("Connected by", addr)
        while True:
            data = conn.recv(1024)
            if not data:
                break
            conn.sendall(data)
```
