package sensitive


import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {

	Load("./words.txt")
	//ww := JSONTree()
	//fmt.Println(ww)
}

func TestHit(t *testing.T) {
	Load("./words.txt")
	sl := map[string]bool{
		"二笔热笔娅": true,
		"热笔娅二笔": true,
		"二笔热笔娅二笔": true,
		"热二笔笔娅": false,
		"二笔热二笔笔娅": false,
		"热二笔笔娅二笔": false,
		"热笔": false,
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
	for s,b := range sl {
		h, r := Hit(s)
		if h != b {
			t.Error(fmt.Sprintf("匹配结果错误: text: %s result: %v", s, r))
		}
	}
}