package main

import "fmt"

// Comparer 定义比较接口
type Comparer interface {
	Compare(data interface{}) bool
}

// LinkedElement 链表的节点
type LinkedElement struct {
	data Comparer
	next *LinkedElement
}

// LinkedList 单链表
type LinkedList struct {
	head   *LinkedElement
	tail   *LinkedElement
	length int
}

// NewList 创建一个新列表
func NewList() *LinkedList {
	head := &LinkedElement{}
	return &LinkedList{
		head: head,
		tail: head,
	}
}

// Insert 指定位置插入元素
func (l *LinkedList) Insert(i int, data Comparer) {

}

// Add 将元素添加到链表中
func (l *LinkedList) Add(data Comparer) {
	el := &LinkedElement{data: data}
	l.tail.next = el
	l.tail = el
	l.length++
}

// Delete 将元素添加到链表中
func (l *LinkedList) Delete(data Comparer) (bool, Comparer) {
	pre := l.head
	p := l.head.next
	for p != nil {
		if p.data.Compare(data) {
			pre.next = p.next
			l.length--
			return true, data
		}
		pre = p
		p = p.next
	}
	return false, nil
}

// DeleteByIndex 将元素添加到链表中
func (l *LinkedList) DeleteByIndex(i int) (bool, Comparer) {
	if i < 1 || i > l.length {
		return false, nil
	}
	pre := l.head
	p := l.head.next
	for j := 1; i > j; j++ {
		pre = p
		p = p.next
	}
	pre.next = p.next
	return true, p.data
}

// Empty 判断列表是否为空
func (l *LinkedList) Empty() bool {
	return l.head.next == nil
}

// Len 判断列表是否为空
func (l *LinkedList) Len() int {
	return l.length
}

// Print 将元素添加到链表中
func (l *LinkedList) Print() {
	p := l.head.next
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
