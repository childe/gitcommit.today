---

layout: post
title: install systemtap
date: 2017-12-20 17:09:02 +0800

---

[http://r00tkit.me/?p=59](http://r00tkit.me/?p=59)

```
	#!/bin/bash
WEB="http://debuginfo.centos.org/6/i386/"
RELEASE=`uname -r`
MACHINE=`uname -m`
PKG1="kernel-debuginfo-$RELEASE.rpm"
PKG2="kernel-debuginfo-common-$MACHINE-$RELEASE.rpm"
wget $WEB$PKG1
wget $WEB$PKG2
#Build Downloaded debuginfo packages
rpm -Uhv kernel-debuginfo-*.rpm
#Install systemtap and kernel-developemnt packages
yum install systemtap kernel-devel
```
