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

func TestIntersectWith(t *testing.T) {
	var s IntSet
	s.Add(2)
	s.Add(3)
	s.Add(5)
	s.Add(7)
	s.Add(11)

	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(4)
	x.Add(5)

	s.IntersectWith(&x)

	if s.String() != "{2 3 5}" {
		t.Fail()
	}
}

func TestDifferenceWith(t *testing.T) {
	var s IntSet
	s.Add(1)
	s.Add(79)
	s.Add(80)
	s.Add(81)
	s.Add(873)
	s.Add(7032)
	s.Add(83840)

	var x IntSet
	x.Add(79)
	x.Add(873)
	x.Add(83840)

	s.DifferenceWith(&x)

	if s.String() != "{1 80 81 7032}" {
		t.Fail()
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	var s IntSet
	s.Add(2)
	s.Add(3)
	s.Add(5)
	s.Add(7)
	s.Add(11)

	var x IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
	x.Add(4)
	x.Add(5)

	s.SymmetricDifferenceWith(&x)

	if s.String() != "{1 4 7 11}" {
		t.Fail()
	}
}
