---

date: 2024-02-18T16:58:34+0800
layout: post

---

问：

openssl 怎么查看一个 bundle 证书的所有证书信息？

答：

要查看一个包含多个证书的证书包（bundle）的所有证书信息，您可以使用OpenSSL命令行工具。首先，确保您已经安装了OpenSSL。

接下来，使用以下命令查看证书包中的所有证书信息：

```
openssl crl2pkcs7 -nocrl -certfile bundle.crt | openssl pkcs7 -print_certs -text -noout
```

在这里，`bundle.crt` 是您的证书包文件名。请将其替换为您的实际文件名。

这个命令首先将证书包转换为PKCS #7格式，然后使用`openssl pkcs7`命令提取和显示所有证书的信息。

如果您只想查看证书包中的某个特定证书，可以使用以下命令：

```
openssl x509 -in bundle.crt -text -noout
```

这将显示证书包中第一个证书的信息。要查看其他证书，请使用`-certopt`选项指定要查看的证书，如下所示：

```
openssl x509 -in bundle.crt -text -noout -certopt ca_only
```

这将显示证书包中的CA证书信息。您可以根据需要调整`-certopt`选项。
