---

date: 2023-02-09T16:41:57+0800
layout: post

---

报错：

```
load program: permission denied: Unrecognized arg#0 type PTR
...
; bpf_map_update_elem(&birth, &pid, &v, BPF_ANY);
20: (18) r1 = 0xffff89a90e035c00
22: (bf) r2 = r7
23: (b7) r4 = 0
24: (85) call bpf_map_update_elem#2
invalid indirect read from stack R3 off -32+12 size 16
processed 23 insns (limit 1000000) max_states_per_insn 0 total_states 1 peak_states 1 mark_read 1
```

原因：struct v 未对齐，需要补充一下 padding
