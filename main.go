package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (node *Node) GetData() string {
	return fmt.Sprintf("%v", node.Data)
}

func preOrderTraverse(treeNode *Node) {
	if treeNode == nil {
		return
	}

	fmt.Printf("%s", treeNode.GetData())
	preOrderTraverse(treeNode.Left)
	preOrderTraverse(treeNode.Right)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	fmt.Printf("header content: [%v]\n", r.Header)
	fmt.Printf("form content:[%s, %s]\n", r.PostFormValue("username"), r.PostFormValue("passwd"))
	bodyContent, err := io.ReadAll(r.Body)
	if err != err && err != io.EOF {
		fmt.Printf("read body content failed, err:[%s]\n", err.Error())
	}
	fmt.Printf("body content:[%s]\n", string(bodyContent))

	fmt.Fprintf(w, "hello world")

}
func main() {
	//
	//file, err := os.Create("aaa.txt")
	//if err != nil {
	//	fmt.Println("Error creating file:", err)
	//	return
	//}
	//defer file.Close()
	//
	//message := "hello, %s!"
	//name := "Gopher"
	//
	//_, err = fmt.Fprintf(file, message, name)
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}
	//
	//fmt.Println("Data written to file successfully")

	http.HandleFunc("/index", IndexHandler)
	http.ListenAndServe("0.0.0.0:8888", nil)

	//node1 := NewNode(0)
	//node2 := NewNode(1)
	//node3 := NewNode(2)
	//node1.Left = node2
	//node2.Right = node3
	//preOrderTraverse(node1)
	//cache := NewLRUCache(2)
	//cache.Put("1", "one")
	//cache.Put("2", "two")
	//cache.Put("3", "three")
	//val, found := cache.Get("2")
	//if found {
	//	fmt.Println(val)
	//}
	now := time.Now()
	fmt.Printf("%T", now)
	fmt.Printf("%v", now)
	print()
}
