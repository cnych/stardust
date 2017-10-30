## 加载toml文件

> 演示加载`toml`文件示例

直接执行：
```shell
go run loader_toml.go
```

得到结果：(失败)
```shell
ERRO[2017-10-30T14:21:29+08:00] Load Config toml                              error="open test.toml: no such file or directory" toml=test.toml
```

得到结果：(成功)
```shell
DEBU[2017-10-30T14:26:06+08:00] Load Config toml                              result="{Site:www.baidu.com}" toml=test.toml
```