package container

import (
	"testing"
)

type KeyInt int

func (k KeyInt) Less(x Comparer) bool {
	return k < x.(KeyInt)
}

func TestInt(t *testing.T) {
	var a KeyInt = 10
	var l Comparer = a
	c := l.Less(KeyInt(11))
	if c != true {
		t.Error("compare error")
	}
}

type KeyString string

func (k KeyString) Less(x Comparer) bool {
	return k < x.(KeyString)
}

func TestString(t *testing.T) {
	var x KeyString = "aaaa"
	var l Comparer = x
	c := l.Less(KeyString("abaa"))
	if c != true {
		t.Error("compare error")
	}
}
