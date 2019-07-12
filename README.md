# stardust

[![Build Status](https://travis-ci.org/cnych/starjazz.svg?branch=master)](https://travis-ci.org/cnych/starjazz)

stardust for go, this is a tooset for golang in daily development~

## 自动生成代码
利用`go generate`特性可以自动生成代码，如执行下面脚本调用`python`脚本生成`useragent`列表：
```shell
source gen_useragent.sh
```

## 测试
可以安装`gotests`插件自动生成测试代码
```shell
go get -u -v github.com/cweill/gotests/...
```
