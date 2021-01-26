// +build OMIT

package main

import (
	"fmt"
	"strings"
)

type List struct {
	head interface{}
	tail *List
}

func New(a ...interface{}) *List {
	var l *List
	for i := len(a) - 1; i >= 0; i-- {
		l = l.Add(a[i])
	}
	return l
}

func NewFromSlice(s []interface{}) *List {
	var l *List
	for i := len(s) - 1; i >= 0; i-- {
		l = l.Add(s[i])
	}
	return l
}



func (list *List) IsEmpty() bool {
	return list == nil
}

func (list *List) Add(a ...interface{}) *List {
	for i := len(a) - 1; i >= 0; i-- {
		list = &List{a[i], list}
	}
	return list
}

func (list *List) AddFromSlice(s []interface{}) *List {
	for i := len(s) - 1; i >= 0; i-- {
		list = &List{s[i], list}
	}
	return list
}

func (list *List) Head() interface{} {
	if list.IsEmpty() { return nil }
	return list.head
}

func (list *List) Tail() *List {
	if list.IsEmpty() { return nil }
	return list.tail
}

func (list *List) String() string {
	var result string
	for i := list; i != nil ; i = i.tail {
		result = result + fmt.Sprint(i.head) + " "
	}
	return "(" + strings.TrimSpace(result) + ")"
}

func main() {
	// create an empty list
	l1 := New()

	// create a list from arguments
	l2 := New(1, 2, 3)

	// create a list from a slice
	slice := []interface{}{1, 2, 3}
	l3 := NewFromSlice(slice)

	fmt.Println(l1, l2, l3)

	l4 := New()

	// add elemnts from arguments
	l5 := l4.Add(1, 2, 3)

	// add elements from a slice
	l6 := l4.AddFromSlice(slice)

	fmt.Println(l5, l6)
}
