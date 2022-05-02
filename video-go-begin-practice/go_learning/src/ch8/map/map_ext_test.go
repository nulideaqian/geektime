package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if _, ok := mySet[1]; ok {
		t.Log("elem is existed")
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	t.Log(len(mySet))
}
