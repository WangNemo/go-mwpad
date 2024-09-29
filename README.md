
<h1 align="center">
  github.com/WangNemo/go-mwpad
</h1>


## 🚀 1. go-mwpad 是什么？
解决golang中对于包含中文字符场景下打印对齐padding的**宽度**问题
即：multiple-width padding

## 🧐 2.为什么用 go-mwpad？

中文字符打印场景下可能会遇到下面的问题：

- **门店名称:&nbsp;Apple&nbsp;王府井&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|取货状态:不可取货**

- **门店名称:&nbsp;Apple&nbsp;三里屯&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|取货状态:不可取货**

- **门店名称:&nbsp;Apple&nbsp;西单大悦城&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|取货状态:不可取货**

- **门店名称:&nbsp;Apple&nbsp;天津恒隆广场&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;|取货状态:不可取货**

`这些"|"为什么对不齐呢？……明明用Printf("%-20s",xxx)这样了`
  

 实际上因为go unicode中的【字符宽度】和【字符长度】是不同的概念，汉字字符的宽度和[East Asian](https://www.google.com/search?q=EAST+ASIAN+WIDTH)规定有关
 
 go官方package里刚好没有，所以简单封装下函数


## 🛠️ 3.如何使用？

__1. 获取package__

```bash
go get github.com/WangNemo/go-mwpad
```

__2. 使用__

```golang
package main

import (
	"fmt"
	mwpad "github.com/WangNemo/go-mwpad"
)

func main() {
	s :="门店名称: Apple 王府井"
	ps, err := mwpad.StrPad(s, 20, nil)
	if err!= nil {
	    fmt.Printf("pad error %s", err.Error())
	}
	fmt.Printf("%s\n", ps)
}

```
__3.支持模式:__
- 从左补齐
- 从右补齐
- 从双边补齐（左侧优先）
- 从双边补齐（右侧优先）
