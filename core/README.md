# ![Version](https://img.shields.io/badge/version-0.0.9-green.svg)

**项目规范**

所有项目目录均放在$GOPATH/src下,目录结构跟gitlab路径一致,例如: `https://gitlab.dev.daaokeji.com/infrastructure/core`, 目录结构如下:
```
$GOPATH
  - src
    - github.cecil777.com
      - infrastructure
        - core
          - 项目文件
```

**编码规范**

- import

对import的包进行分组管理，用换行符分割，而且标准库作为分组的第一组。如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，建议采用如下方式进行组织你的包
```
package main

import (
    "fmt"
    "os"

    "kmg/a"
    "kmg/b"

    "code.google.com/a"
    "github.com/b"
)
```

- go文件命名

go文件名均为小写,单词与单词之间用`-`分割, 例如: `i-string-generator.go`

- 包名

包名为小写,单词与单词直接拼接,例如: `userservice`

- go文件内容

```
package mypackage

import (
    公用包

    内部包

    第三方包
)

const (
    PublicConstA = '开放常量A'
    PublicConstB = '开放常量B'

    privateConstA = '私有常量A'
    privateConstB = '私有常量B'
)

var (
    PublicVariableA = '开放变量A'
    PublicVariableB = '开放变量B'

    privateVariableA = '私有变量A'
    privateVariableB = '私有变量B'
)

type myStruct struct {
    SuperStruct

    PublicA string
    PublicB int

    privateA string
    privateB int
}

func (m myStruct) PublicMethodA() {}

func (m myStruct) PublicMethodB() {}

func (m myStruct) privateMethodA() {}

func (m myStruct) privateMethodB() {}

func PublicMethodA() {}

func PublicMethodB() {}

func init() {}

func privateMethodA() {}

func priateMethodB() {}
```