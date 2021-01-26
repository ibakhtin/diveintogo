package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Node struct {
	head interface{}
	tail *Node
}

type List struct {
	data *Node
	len  int
}

func New() *List {
	return new(List)
}

func (l *List) Head() (interface{}, bool) {
	if l.len > 0 {
		return l.data.head, true
	}
	return nil, false
}

func (l *List) Tail() (*List, bool) {
	if l.len > 0 {
		return &List{l.data.tail, l.len - 1}, true
	}
	return nil, false
}

func (l *List) Add(v interface{}) *List {
	l.data = &Node{v, l.data}
	l.len += 1
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) String() string {
	s := "("
	for node := l.data; node != nil; node = node.tail {
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