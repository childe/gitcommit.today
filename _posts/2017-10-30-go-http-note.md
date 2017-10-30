---

layout: post
title: golang http note
date: 2017-10-30 10:03:57 +0800

---

没看懂这是啥, 先继续往下看.

  TrailerPrefix is a magic prefix for ResponseWriter.Header map keys that, if present, signals that the map entry is actually for the response trailers, and not the response headers. The prefix is stripped after the ServeHTTP call finishes and the values are sent in the trailers.

  This mechanism is intended only for trailers that are not known prior to the headers being written. If the set of trailers is fixed or known before the header is written, the normal Go trailers mechanism is preferred:

  https://golang.org/pkg/net/http/#ResponseWriter
  https://golang.org/pkg/net/http/#example_ResponseWriter_trailers
  const TrailerPrefix = "Trailer:"


MDN web docs里面的解释 [https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Trailer](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Trailer)
