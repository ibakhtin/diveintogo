// +build OMIT

package main

import (
	"fmt"
	"strings"
)

type LinkedList interface{
	New() *LinkedList
}

type Value interface{}

type List struct {
	head Value
	tail *List
}

func New() *List {
	return nil
}

func (l *List) IsEmpty() bool {
	return l == nil
}

func (l *List) Add(v Value) *List {
	return &List{v, l}
}

func (l *List) Head() Value {
	if l.IsEmpty() {
		return nil
	} else {
		return l.head
	}
}

func (l *List) Tail() *List {
	if l.IsEmpty() {
		return nil
	} else {
		return l.tail
	}
}

func (l *List) String() string {
	result := "("
	for h := l; h != nil ; h = h.tail {
		result = result + fmt.Sprint(h.head) + " "
	}
	return strings.TrimSpace(result) + ")"
}

func main() {
	a := New().Add(1)
	b := a.Add(2)
	fmt.Println(a, b)
}
