package intset

import "testing"

func TestLen(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(100)
	s.Add(1024)
	s.Add(65530)

	if s.Len() != 5 {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(100)
	s.Add(1024)
	s.Add(65530)

	s.Remove(100)
	s.Remove(1024)

	if s.String() != "{1 5 65530}" {
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(100)
	s.Add(1024)
	s.Add(65530)

	s.Clear()
	if s.String() != "{}" {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(5)
	s.Add(100)
	s.Add(1024)
	s.Add(65530)

	c := s.Copy()
	c.Add(2)
	c.Remove(65530)

	if s.String() != "{1 5 100 1024 65530}" {
		t.Fail()
	}
	if c.String() != "{1 2 5 100 1024}" {
		t.Fail()
	}
}

func TestAddAll(t *testing.T) {
	var s IntSet
	s.AddAll(1, 5, 100, 1024, 65530)

	if s.String() != "{1 5 100 1024 65530}" {
		t.Fail()
	}
}
