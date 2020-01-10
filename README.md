# go-sensitive

敏感词检测

## Usage
```golang
package main

import (
    "fmt"
    "github.com/chenjinya/go-sensitive"
)

func main(){

    
    sensitive.Load("./words.txt")
    h,s := sensitive.Hit("热比娅亚")

    fmt.Println(h, s)
}
```

需要提前准备 `words.txt` 敏感词库

## 逻辑

使用多叉树模型



