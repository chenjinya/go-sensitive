package sensitive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Tree 类型
type Tree map[string]*Node

//存放Tree变量
var tree  = make(Tree)

//Node结构体
type Node struct {
	Text string `json:"text"`
	Nodes Tree `json:"nodes"`
}

//初始化词汇为slice
func formatWords(ws []string) [][]string{
	var l [][]string
	for _, v := range ws {
		s := strings.Split(v, "")
		l = append(l, s)
	}
	return l
}

//建立多叉树结构
func buildTree(wd []string, n map[string]*Node){
	if len(wd) == 0 {
		return
	}
	s := wd[0]
	if n[s] == nil {
		n[s] = &Node{
			Text: s,
			Nodes: nil,//最后节点为空
		}
	}
	if len(wd) > 1 {
		n[s].Nodes = make(map[string]*Node)
		buildTree(wd[1:], n[s].Nodes)
	}

}

//广度优先遍历
func BFSTravelTree(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Text)
	for _, l := range n.Nodes {
		BFSTravelTree(l)
	}
}

//输出JSON结构
func JSONTree() string{
	j, _ :=  json.MarshalIndent(tree, "", "  ")
	return fmt.Sprintf("%s \n",j)
}

//匹配递归
func match(ws []string, t *Tree, rs string) (bool,string){
	if len(ws) == 0 {
		return false,""
	}
	if (*t)[ws[0]] == nil {
		return match(ws[1:],&tree, "")
	}

	if (*t)[ws[0]].Nodes == nil {
		return true, rs + ws[0]
	}
	return match(ws[1:],&(*t)[ws[0]].Nodes,  rs + ws[0])
}

//命中查询
func Hit(w string) (bool, string){
	ws := strings.Split(w, "")
	return match(ws, &tree, "")
}

//初始化
func Load(p string){

	f, err := os.Open(p)
	if err != nil {
		panic("read file fail")
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		panic("read to fd fail")
	}

	s := string(fd);
	ws := formatWords(strings.Split(s, "\n"))
	for _, sl := range ws {
		buildTree(sl, tree)
	}
}

