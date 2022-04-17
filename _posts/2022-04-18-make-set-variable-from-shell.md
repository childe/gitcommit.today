---

layout: post
date: 2022-04-18T00:35:20+0800

---

```
$(eval GID=$(shell docker exec jenkins-slave ls -n /var/run/docker.sock | cut -f 4 -d ' '))
echo $(GID)
docker exec jenkins-slave groupadd -g $(GID) docker
```
