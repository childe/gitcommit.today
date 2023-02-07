---

date: 2023-02-07T17:25:18+0800
title: 'tags in docker image name'
layout: post

---

In docker-Context here are the important informations you need.

Alpine
Url: https://alpinelinux.org/
Imagename: alpine
Shorty: Its very small.
Packagemanger: apk
Shells: /bin/sh
Size: Few MBs - current tag needs 2.7MB

Jessie aka Debian 8
Url: https://wiki.debian.org/DebianJessie
Imagename: debian:jessie
Shorty: No LTS anymore
Packagemanager: apt
Shells: /bin/bash
Size: ~50mb

Stretch aka Debian 9
Url: https://wiki.debian.org/DebianStretch
Imagename: debian:stretch
Shorty: LTS is running out
Packagemanager: apt
Shells: /bin/bash and many more
Size: ~40mb

Buster aka Debian 10
Url: https://wiki.debian.org/DebianBuster
Imagename: debian:buster
Shorty: All what you need, but newer
Packagemanager: apt
Shells: /bin/bash and many more
Size: ~50mb

Bullseye aka Debian 11
Url: https://wiki.debian.org/DebianBullseye
Imagename: debian:bullseye
Shorty: Newest debian
Shells: /bin/bash and many more
Size: ~50mb

Ubuntu based on debain
Url: https://hub.docker.com/_/ubuntu
Imagename: ubuntu
Shorty: All what you need
Packagemanager: apt
Shells: /bin/bash and more
Size: ~25mb


You can find a list of Debian releases and their end of life (EOL) dates here: [wiki.debian.org/DebianReleases](https://wiki.debian.org/DebianReleases)
