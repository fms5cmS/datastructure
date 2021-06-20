package list

import "testing"

func TestList_String(t *testing.T) {
	l := NewList()
	t.Logf("the elements of list: %v", l)
	l.dummyHead.next = &Node{
		Val: 100,
		next: &Node{
			Val: 20,
			next: &Node{
				Val:  -50,
				next: nil,
			},
		},
	}
	t.Logf("the elements of list: %v", l)
}

func TestList_Insert(t *testing.T) {
	list := NewList()
	list.Insert("first")
	list.Insert("second")
	list.Insert(3)
	list.Insert("last")
	t.Logf("list: %v, size: %d", list, list.Size())
}

func TestList_InsertAfter(t *testing.T) {
	list := NewList()
	list.Insert("first")
	t.Logf("list: %v, size: %d", list, list.Size())
	list.Insert("second")
	list.Insert("fourth")
	t.Logf("list: %v, size: %d", list, list.Size())
	if err := list.InsertAfter(2, "third"); err != nil {
		t.Log(err)
	}
	list.InsertAfter("second", "third")
	t.Logf("list: %v, size: %d", list, list.Size())
}

func TestList_Remove(t *testing.T) {
	list := NewList()
	list.Insert(100)
	list.Insert(90)
	list.Insert(95)
	list.Insert(80)
	t.Logf("list: %v, size: %d", list, list.Size())
	if err := list.Remove(77); err != nil {
		t.Log(err)
	}
	list.Remove(95)
	t.Logf("list: %v, size: %d", list, list.Size())
	list.Remove(80)
	t.Logf("list: %v, size: %d", list, list.Size())
}

func TestList_Update(t *testing.T) {
	list := NewList()
	list.Insert(1)
	list.Insert(2)
	list.Insert("third")
	t.Logf("list: %v, size: %d", list, list.Size())
	if err := list.Update(3, "three"); err != nil {
		t.Log(err)
	}
	list.Update("third", 3)
	t.Logf("list: %v, size: %d", list, list.Size())
}
