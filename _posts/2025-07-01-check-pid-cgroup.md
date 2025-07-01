---

date: 2025-07-01T11:11:32+0800
layout: post

---


```sh
# cat /proc/<PID>/cgroup
1:name=systemd:/user.slice/user-1000.slice/session-2.scope
3:cpu,cpuacct:/user.slice
5:memory:/docker/容器ID

# ps --cgroup -p <PID>
PID CGROUP
1234 5:memory:/docker/nginx,3:cpu:/docker/nginx
```
