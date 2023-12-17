package test

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	A = 1 + iota
	B = 3 + iota // 增量为2
	C
)

func TestConstantTry(t *testing.T) {
	t.Log(A, B, C)
}
