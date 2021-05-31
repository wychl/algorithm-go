package lru

import (
	"reflect"
	"testing"
)

func TestLru(t *testing.T) {
	t.Run("Test Push()", func(t *testing.T) {
		// 初始化容量为6的队列
		list := NewList(6)
		list.Push(1)
		list.Push(2)
		list.Push(3)
		want := []interface{}{3, 2, 1}
		got := []interface{}{}
		current := list.Head
		got = append(got, current.Value)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Value)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test Pop()", func(t *testing.T) {
		list := NewList(3)
		list.Push(1)
		list.Push(2)
		list.Push(3)
		list.Pop()
		want := []interface{}{3, 2}
		got := []interface{}{}
		current := list.Head
		got = append(got, current.Value)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Value)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Test View()", func(t *testing.T) {
		list := NewList(6)
		list.Push(1)
		list.Push(2)
		list.Push(3)
		list.View(1)
		want := []interface{}{1, 3, 2}
		got := []interface{}{}
		current := list.Head
		got = append(got, current.Value)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Value)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}
