package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Next *Node
	Prev *Node
	key interface{}
}

func InitializeHead(head *Node) {
	head.Next = head
	head.Prev = head
}

//adds to front
func InsertAtHead(head *Node, node *Node){
	head.Prev = node
	node.Prev = node
	node.Next = head
}
//removes from front
func Remove(head *Node, key interface{}){
	tmp := head.Next
	for tmp != head {
		if tmp.key == key {
			fmt.Println("Found Key")
			tmp.Next.Prev = tmp.Prev
			tmp.Prev.Next = tmp.Next
			break
		}
		tmp = tmp.next
	}
}

func Print(head *Node){
	fmt.Print("{")
	nptr := lst.Head
	for nptr != nil{
		fmt.Print(nptr.Value)
		fmt.Print(", ")
		nptr = nptr.Next
	}
	fmt.Println("}")
}

func main() {
	lst := &llst{nil, nil}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Commands: \"add\", \"remove\", \"print\"")
	for {
		line, _ := reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "add":
			value, _ := strconv.ParseInt(parts[1], 0,64)
			lst.Add((int(value)))
		case "remove":
			lst.Remove()
		default:
			lst.Print()
		}
	}
}
