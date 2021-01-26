// +build OMIT

package main

import (
	"fmt"
)

type List interface {
	Add(interface{}) List
	Head() (interface{}, bool)
	IsEmpty() bool
	Tail() (List, bool)
}

type emptyList struct{}

type list struct {
	head interface{}
	tail List
}

func (e *emptyList) Add(head interface{}) List {
	return &list{head, e}
}

func (l *list) Add(head interface{}) List {
	return &list{head, l}
}


func (e *emptyList) Head() (interface{}, bool) {
	return nil, false
}
func (l *list) Head() (interface{}, bool) {
	return l.head, true
}


func (e *emptyList) IsEmpty() bool {
	return true
}
func (e *list) IsEmpty() bool {
	return false
}


func (e *emptyList) Tail() (List, bool) {
	return nil, false
}
func (l *list) Tail() (List, bool) {
	return l.tail, true
}


func (e *emptyList) String() string {
	return "()"
}
func (l *list) String() string {
	var s string
	curr := l
	for {
		s = s + fmt.Sprint(curr.head) + " "
		tail, _ := curr.Tail()
		if tail.IsEmpty() {
			return "(" + strings.TrimSpace(s) + ")"
		}
		curr = tail.(*list)
	}
}


func New() List {
	return &emptyList{}
}

func main() {
	a := New()
	a = a.Add(2).Add(1)
	fmt.Println(a)
}