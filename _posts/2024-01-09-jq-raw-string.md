---

date: 2024-01-09T15:03:47+0800
layout: post

---

使用 linux 的命令 jq 输出json 对象中的一个字段 name，name是一个字符串。 我使用 jq '.name' 输出的结果带有双引号，有什么办法可以不要双引号吗？最好是 jq 命令自己的参数，而不是使用管道交给另外一个命令处理。


要使用 `jq` 命令输出 JSON 对象中的一个字段（例如 `name`），并去除结果中的双引号，你可以使用 `-r` 或 `--raw-output` 选项。这将输出原始字符串，而不是 JSON 字符串。例如：

```bash
jq -r '.name'
```