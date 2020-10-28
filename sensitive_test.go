package sensitive


import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var example = map[string]bool{
"二笔热笔娅": true,
"热笔娅二笔": true,
"二笔热笔娅二笔": true,
"热二笔笔娅": false,
"二笔热二笔笔娅": false,
"热二笔笔娅二笔": false,
"热笔": true,
"热笔笔笔": true,
"笔娅": false,
`
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
一二三四五六七八九〇
`: false,
}

func BenchmarkHit(b *testing.B){
	Load("./words.txt")
	sl := example

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		for s,_ := range sl {
			Hit(s)
		}
	}
}


//普通的遍历关键词测试
func BenchmarkTraverseHit(b *testing.B){
	f, err := os.Open("./words.txt")
	if err != nil {
		panic("read file fail")
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic("read to fd fail")
	}

	s := string(fd)
	ws := strings.Split(s, "\n")

	sl := example

	b.ResetTimer()
	for i:=0;i<b.N;i++{
		for s,_ := range sl {
			for _,v := range ws {
				if strings.Index(s, v) >= 0 {
					break
				}
			}
		}
	}
}

func TestFile(t *testing.T) {

	Load("./words.txt")
	// ww := JSONTree()
	// fmt.Println(ww)
}

func TestHit(t *testing.T) {
	Load("./words.txt")
	sl := example
	for s,b := range sl {
		h, r := Hit(s)
		if h != b {
			t.Error(fmt.Sprintf("匹配结果错误: text: %s result: %v", s, r))
		}
	}
}
