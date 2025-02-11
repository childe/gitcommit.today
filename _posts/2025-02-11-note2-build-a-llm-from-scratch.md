---

date: 2025-02-11T14:18:49+0800
layout: post

---

需要看pytorch的文档，进一步了解里面第一步的作用是什么，以了解MultiHeadAttention的原理

```
class MultiHeadAttention(nn.Module):
    def __init__(self, d_in, d_out, context_length, dropout, num_heads, qkv_bias=False):
        super().__init__()
        assert d_out % num_heads == 0, "d_out must be divisible by num_heads"

        self.d_out = d_out
        self.num_heads = num_heads
        self.head_dim = d_out // num_heads  # 1
        self.W_query = nn.Linear(d_in, d_out, bias=qkv_bias)
        self.W_key = nn.Linear(d_in, d_out, bias=qkv_bias)
        self.W_value = nn.Linear(d_in, d_out, bias=qkv_bias)
        self.out_proj = nn.Linear(d_out, d_out)  # 2
        self.dropout = nn.Dropout(dropout)
        self.register_buffer(
            "mask", torch.triu(torch.ones(context_length, context_length), diagonal=1)
        )

    def forward(self, x):
        b, num_tokens, d_in = x.shape
        keys = self.W_key(x)  # 3
        queries = self.W_query(x)  # 3
        values = self.W_value(x)  # 3

        keys = keys.view(b, num_tokens, self.num_heads, self.head_dim)  # 4
        values = values.view(b, num_tokens, self.num_heads, self.head_dim)
        queries = queries.view(b, num_tokens, self.num_heads, self.head_dim)

        keys = keys.transpose(1, 2)  # 5
        queries = queries.transpose(1, 2)  # 5
        values = values.transpose(1, 2)  # 5

        attn_scores = queries @ keys.transpose(2, 3)  # 6
        mask_bool = self.mask.bool()[:num_tokens, :num_tokens]  # 7

        attn_scores.masked_fill_(mask_bool, -torch.inf)  # 8

        attn_weights = torch.softmax(attn_scores / keys.shape[-1] ** 0.5, dim=-1)
        attn_weights = self.dropout(attn_weights)

        context_vec = (attn_weights @ values).transpose(1, 2)  # 9
        # 10
        context_vec = context_vec.contiguous().view(b, num_tokens, self.d_out)
        context_vec = self.out_proj(context_vec)  # 11
        return context_vec
```
