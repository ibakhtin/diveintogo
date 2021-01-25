package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Value interface{}

type Node struct {
	head Value
	tail *Node
}

type List struct {
	node *Node
	len  int
}

func New() *List {
	return new(List)
}

func (l *List) Head() (Value, bool) {
	if l.len > 0 {
		return l.node.head, true
	}
	return nil, false
}

func (l *List) Tail() (*List, bool) {
	if l.len > 0 {
		return &List{l.node.tail, l.len - 1}, true
	}
	return nil, false
}

func (l *List) Add(v Value) *List {
	l.node = &Node{v, l.node}
	l.len += 1
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) String() string {
	s := "("
	for node := l.node; node != nil; node = node.tail {
		if reflect.TypeOf(node.head).String() == "string" {
			s += fmt.Sprintf("%q", node.head) + " "
		} else {
			s += fmt.Sprint(node.head) + " "
		}
	}
	return strings.TrimSpace(s) + ")"
}

func logFatal(e error) {
	log.Fatal(e)
}

func main() {
	linkedList := New()

	head, ok := linkedList.Head()

	if ok {
		fmt.Println(head, linkedList.Len())
	} else {
		fmt.Println("List is empty.")
	}

	linkedList.Add("i like golang")

	head, ok = linkedList.Head()

	if ok {
		fmt.Printf("%v %T %d\n", head, head, linkedList.Len())
	} else {
		fmt.Println("List is empty.")
	}
	linkedList.Add(42)
	fmt.Println(linkedList)
}