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
    h,s := sensitive.Hit("热比垭亚")

    fmt.Println(h, s)
}
```

需要提前准备 `words.txt` 敏感词库

## 逻辑

使用多叉树模型

## Benchmark

name |循环次数 |  执行时间 | 占用内存 | 分配内存次数 
-------- | -----| -----| ----- | -----
50个关键词| 
BenchmarkHit-4  |                  124059   |          10298 ns/op      |      2480 B/op    |     12 allocs/op |
BenchmarkTraverseHit-4       |     123309    |          9604 ns/op      |         0 B/op     |     0 allocs/op |
100个关键词|
BenchmarkHit-4               |     110617    |          9628 ns/op       |     2480 B/op     |    12 allocs/op |
BenchmarkTraverseHit-4       |      58798    |         19927 ns/op       |        0 B/op     |     0 allocs/op |

循环次数 |  50 多叉树    | 50 遍历        | 100 多叉树 | 100 遍历 
--------| -----        | -----         | -----      | -----
循环次数 | 124059       | 123309        | 110617        | 58798
执行时间 | 10298 ns/op  | 9604 ns/op    | 9628 ns/op    | 19927 ns/op
占用内存 | 2480 B/op    | 0 B/op        | 2480 B/op     |  0 B/op
分配内存 | 12 allocs/op | 0 allocs/op   | 12 allocs/op  | 0 allocs/op

在敏感词 大于 50 个的时候，多叉树模型比传统遍历性能有显著提升

