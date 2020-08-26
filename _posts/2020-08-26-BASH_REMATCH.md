---

date: 2020-08-26T23:32:46+0800

---

[Bash Regular Expressions](https://www.linuxjournal.com/content/bash-regular-expressions)

```shell
#!/bin.bash

if [[ $# -lt 2 ]]; then
    echo "Usage: $0 PATTERN STRINGS..."
    exit 1
fi
regex=$1
shift
echo "regex: $regex"
echo

while [[ $1 ]]
do
    if [[ $1 =~ $regex ]]; then
        echo "$1 matches"
        i=1
        n=${#BASH_REMATCH[*]}
        while [[ $i -lt $n ]]
        do
            echo "  capture[$i]: ${BASH_REMATCH[$i]}"
            let i++
        done
    else
        echo "$1 does not match"
    fi
    shift
done
```

```
  # sh bashre.sh 'aa(b{2,3}[xyz])cc' aabbxcc aabbcc
  regex: aa(b{2,3}[xyz])cc

  aabbxcc matches
    capture[1]: bbx
  aabbcc does not match
```
