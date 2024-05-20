---

date: 2024-04-02T20:33:05+0800
layout: post

---

遇到下面这样的报错

```
  opensslErrorStack: [ 'error:03000086:digital envelope routines::initialization error' ],
  library: 'digital envelope routines',
  reason: 'unsupported',
  code: 'ERR_OSSL_EVP_UNSUPPORTED'
```

export NODE_OPTIONS=--openssl-legacy-provider
