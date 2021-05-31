package lru

import (
	"fmt"
	"strings"
)

// Node 节点规格
type Node struct {
	Value interface{}
	Next  *Node
}

// NewNode 生成新节点
func NewNode(val interface{}) *Node {
	return &Node{val, nil}
}

type List struct {
	Head     *Node
	Length   int
	Capacity int
}

// NewList 初始化队列
func NewList(capacity int) *List {
	return &List{
		Capacity: capacity,
	}
}

// Append 添加值到队列首部
func (l *List) Push(val interface{}) {
	if l.Length >= l.Capacity {
		l.Pop()
	}

	newHead := NewNode(val)
	if l.Head == nil {
		l.Head = newHead
	} else {
		newHead.Next = l.Head
		oldHead := l.Head
		newHead.Next = oldHead
		l.Head = newHead
	}

	l.Length++
}

// Pop 出队列
func (l *List) Pop() *Node {
	cue := l.Head
	var pre *Node

	for cue.Next != nil {
		pre = cue
		cue = cue.Next
	}
	pre.Next = nil

	return cue
}

// View 访问数据
func (l *List) View(val interface{}) {
	var pre *Node

	cue := l.Head

	for cue != nil {
		if cue.Value == val {
			pre.Next = cue.Next
			cue.Next = l.Head
			l.Head = cue
			return
		}

		pre = cue
		cue = cue.Next
	}
}

// Travel 遍历
func (l *List) Travel() string {
	cue := l.Head
	var r []string
	for cue != nil {
		r = append(r, fmt.Sprintf("%v", cue.Value))
		cue = cue.Next
	}
	return strings.Join(r, "-")
}

// Clear 清除
func (l *List) Clear() {
	l.Length = 0
	l.Head = nil
}
