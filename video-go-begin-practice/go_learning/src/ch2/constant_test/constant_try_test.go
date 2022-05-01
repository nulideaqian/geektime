package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTest(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writable, Executable)
}
