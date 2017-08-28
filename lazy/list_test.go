package lazyskiplist

import (
	"github.com/zoyi/skiplist/lib"
	"reflect"
	"testing"
)

func TestPut(t *testing.T) {
	list := NewLazySkipList(lib.IntComparator)
	list.Put(1, "test")
	value, _ := list.Get(1)
	if value != "test" {
		t.Errorf("Expected: %v, Got: %v", "test", value)
	}
}

func TestForwardIterator(t *testing.T) {
	list := NewLazySkipList(lib.IntComparator)

	list.Put(1, nil)
	list.Put(3, nil)
	list.Put(5, nil)
	list.Put(7, nil)

	var slice []int
	for it := list.Begin(); it.Present(); it.Next() {
		slice = append(slice, it.Key().(int))
	}
	expected := []int{1, 3, 5, 7}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, slice)
	}
}

func TestBackwardIterator(t *testing.T) {
	list := NewLazySkipList(lib.IntComparator)

	list.Put(1, nil)
	list.Put(3, nil)
	list.Put(5, nil)
	list.Put(7, nil)

	var slice []int
	for it := list.End(); it.Present(); it.Prev() {
		slice = append(slice, it.Key().(int))
	}
	expected := []int{7, 5, 3, 1}
	if !reflect.DeepEqual(slice, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, slice)
	}
}
