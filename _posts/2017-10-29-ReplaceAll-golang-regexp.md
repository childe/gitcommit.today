---

layout: post
title:  "ReplaceAll in go"
date:   2017-10-29 10:51:15 +0800

---

go文档的[https://golang.org/pkg/regexp/#Regexp.ReplaceAll](https://golang.org/pkg/regexp/#Regexp.ReplaceAll)里面说:

  func (*Regexp) ReplaceAll
  func (re *Regexp) ReplaceAll(src, repl []byte) []byte
  ReplaceAll returns a copy of src, replacing matches of the Regexp with the replacement text repl.
  Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.

```
p, _ = regexp.Compile("(a+)")
b = p.ReplaceAll([]byte("baabaaa"), []byte(`$1`))
glog.Infof("%s\n", b)
b = p.ReplaceAll([]byte("baabaaa"), []byte("xyz"))
glog.Infof("%s\n", b)
```

结果如下:

```
baabaaa
bxyzbxyz
```

Split也和我想像的不一样

```
s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
// s: ["", "b", "b", "c", "cadaaae"]
```
