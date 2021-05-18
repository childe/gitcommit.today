---

date: 2021-05-18T16:37:30+0800

---

从 8.8.8.8 返回了多个 UDP 包。看起来，第一个返回时间太短，TTL 也不对。是不是被污染了？

```
[root@VMS160953 ~]# tcpdump -nn -vvvv host 8.8.8.8
tcpdump: listening on eth0, link-type EN10MB (Ethernet), capture size 262144 bytes
16:36:18.031876 IP (tos 0x0, ttl 64, id 25326, offset 0, flags [none], proto UDP (17), length 126)
    10.60.225.127.47457 > 8.8.8.8.53: [bad udp cksum 0xfc46 -> 0x6d26!] 55814+ [1au] A? prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. ar: . OPT UDPsize=4096 (98)
16:36:18.035181 IP (tos 0x0, ttl 35, id 48898, offset 0, flags [DF], proto UDP (17), length 131)
    8.8.8.8.53 > 10.60.225.127.47457: [udp sum ok] 55814 q: A? prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. 1/0/0 prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. [3m16s] A 104.244.46.186 (103)
16:36:18.035199 IP (tos 0x0, ttl 29, id 0, offset 0, flags [none], proto UDP (17), length 131)
    8.8.8.8.53 > 10.60.225.127.47457: [udp sum ok] 55814 q: A? prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. 1/0/0 prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. [1m51s] A 4.78.139.54 (103)
16:36:18.271586 IP (tos 0x0, ttl 94, id 28433, offset 0, flags [none], proto UDP (17), length 163)
    8.8.8.8.53 > 10.60.225.127.47457: [udp sum ok] 55814 q: A? prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. 2/0/1 prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. [33s] CNAME s3-r-w.eu-central-1.amazonaws.com., s3-r-w.eu-central-1.amazonaws.com. [1s] A 52.219.75.128 ar: . OPT UDPsize=512 (135)
16:36:18.271618 IP (tos 0xc0, ttl 64, id 47249, offset 0, flags [none], proto ICMP (1), length 191)
    10.60.225.127 > 8.8.8.8: ICMP 10.60.225.127 udp port 47457 unreachable, length 171
	IP (tos 0x0, ttl 94, id 28433, offset 0, flags [none], proto UDP (17), length 163)
    8.8.8.8.53 > 10.60.225.127.47457: [udp sum ok] 55814 q: A? prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. 2/0/1 prod-eu-central-1-starport-layer-bucket.s3.eu-central-1.amazonaws.com. [33s] CNAME s3-r-w.eu-central-1.amazonaws.com., s3-r-w.eu-central-1.amazonaws.com. [1s] A 52.219.75.128 ar: . OPT UDPsize=512 (135)
^C
5 packets captured
5 packets received by filter
0 packets dropped by kernel
[root@VMS160953 ~]# ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_seq=1 ttl=95 time=241 ms
^C
--- 8.8.8.8 ping statistics ---
1 packets transmitted, 1 received, 0% packet loss, time 0ms
rtt min/avg/max/mdev = 241.706/241.706/241.706/0.000 ms
```
