package string_test

import "testing"

func TestUnicode(t *testing.T) {
	s := "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c), c[0], s)
}
