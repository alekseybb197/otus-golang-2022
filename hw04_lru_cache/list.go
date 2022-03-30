package hw04lrucache

import (
	"fmt"
	"reflect"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	fmt.Stringer
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	FrontItem *ListItem
	BackItem  *ListItem
	Length    int
}

func (p *list) Len() int         { return p.Length }
func (p *list) Front() *ListItem { return p.FrontItem }
func (p *list) Back() *ListItem  { return p.BackItem }
func (p *list) PushFront(v interface{}) *ListItem {
	newitem := &ListItem{Value: v, Next: p.FrontItem, Prev: nil} // new list item
	if p.FrontItem == nil && p.BackItem == nil {
		p.FrontItem = newitem
		p.BackItem = newitem
	} else {
		if p.FrontItem != nil {
			p.FrontItem.Prev = newitem // fix old head
		}
		p.FrontItem = newitem
	}
	p.Length++
	return newitem
}
func (p *list) PushBack(v interface{}) *ListItem {
	newitem := &ListItem{Value: v, Next: nil, Prev: p.BackItem} // new list item
	if p.FrontItem == nil && p.BackItem == nil {
		p.FrontItem = newitem
		p.BackItem = newitem
	} else {
		if p.BackItem != nil {
			p.BackItem.Next = newitem // fix old tail
		}
		p.BackItem = newitem
	}
	p.Length++
	return newitem
}
func (p *list) Remove(i *ListItem) {
	if i == nil || p.Length == 0 || reflect.TypeOf(*i).Name() != "ListItem" {
		return
	}
	prev := i.Prev   // item forward
	next := i.Next   // item backward
	if next != nil { // if this item is not tail
		next.Prev = prev
	} else {
		if prev != nil {
			prev.Next = nil
		}
		p.BackItem = prev
	}
	if prev != nil { // if this item is not head
		prev.Next = next
	} else {
		if next != nil {
			next.Prev = nil
		}
		p.FrontItem = next
	}
	p.Length--
}
func (p *list) MoveToFront(i *ListItem) {
	if i != nil && p.Length != 0 && reflect.TypeOf(*i).Name() == "ListItem" {
		v := i.Value
		p.Remove(i)
		_ = p.PushFront(v)
	}
}
func (p *list) String() string { // so, but how can I debug this?
	s := "["
	pp := p.FrontItem
	for {
		if pp == nil {
			break
		}
		s += fmt.Sprintf(" %v", pp.Value)
		pp = pp.Next
	}
	s += " ]"
	return s
}

func NewList() List {
	return new(list)
}
