---

date: 2022-04-24T13:25:45+0800
layout: post

---

> https://itsiti.com/how-to-determine-linux-os-endian-big-endian-little-endian/

## Command 1

```sh
echo I | tr -d [:space:] | od -to2 | awk 'FNR==1{ print substr($2,6,1)}'
```

## Command 2

```sh
printf '\1' | od -dAn
```

## Command 3

```sh
lscpu | grep Endian
```

## Command 4

```sh
echo -n I | od -to2 | awk '{ print substr($2,6,1); exit}'
```

## Command 5

```sh
python -c "import sys; print(sys.byteorder)"
```
