package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type llst struct {
	Head *node
	Tail *node
}

type node struct {
	Next *node
	Prev *node
	Value int
}

//adds to back
func (lst *llst) add(value int){
	if lst.Head == nil{
		next := &node{nil, nil, value}
		lst.Head = next
		lst.Tail = next
	} else{
		next := node{nil, lst.Tail, value}
		lst.Tail.Next = &next
		lst.Tail = &next
	}
}
//removes from front
func (lst *llst) remove(){
	if lst.Head != nil{
		lst.Head = lst.Head.Next
	}
}

func (lst *llst) print(){
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
			lst.add((int(value)))
		case "remove":
			lst.remove()
		default:
			lst.print()
		}
	}
}
