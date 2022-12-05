package stack

import "testing"

func TestPush(t *testing.T) {
	s := New[interface{}](nil)
	s.Push(1)
	if s.Peek() != 1 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	s := New[interface{}](nil)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Pop() != 3 {
		t.Fail()
	}
}

func TestPushN(t *testing.T) {
	s := New[interface{}](nil)
	s.PushN([]interface{}{1, 2, 3})
	if s.Pop() != 3 {
		t.Fail()
	}
}

func TestPopN(t *testing.T) {
	s := New[interface{}](nil)
	s.PushN([]interface{}{1, 2, 3})
	if s.PopN(2)[0] != 2 {
		t.Fail()
	}
}

func TestPeek(t *testing.T) {
	s := New[interface{}](nil)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Peek() != 3 {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	s := New[interface{}](nil)
	s.PushN([]interface{}{1, 2, 3})
	s.Reverse()
	if s.Pop() != 1 {
		t.Fail()
	}
}

func TestPushLeft(t *testing.T) {
	s := New[interface{}](nil)
	s.PushLeft(1)
	s.PushLeft(2)
	s.PushLeft(3)
	if s.Pop() != 1 {
		t.Fail()
	}
}

func TestPopLeft(t *testing.T) {
	s := New[interface{}](nil)
	s.PushN([]interface{}{1, 2, 3})
	if s.PopLeft() != 1 {
		t.Fail()
	}
}
