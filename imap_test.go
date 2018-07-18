package imap

import (
	"testing"
)

func TestMap(t *testing.T) {
	t.Helper()

	f := func(key interface{}) {
		m := NewMap()
		m.Add(key, 1)
		m.Add(key, 1.0)
		m.Add(key, "aaabbb")
		if _, ok := m.Get(key); ok == true {
			t.Fatal("should be not ok")
		}
		m.Apply()
		m.Add(key, "aaa")
		val, ok := m.Get(key)
		if ok == false || val.(string) != "aaabbb" {
			t.Fatal("should be not false")
		}
	}

	f(1)
	f(1.0)
	f("1")
}

func TestMapI(t *testing.T) {
	t.Helper()

	f := func(key int) {
		m := NewMapI()
		m.Add(key, 1)
		m.Add(key, 1.0)
		m.Add(key, "aaabbb")
		if _, ok := m.Get(key); ok == true {
			t.Fatal("should be not ok")
		}
		m.Apply()
		m.Add(key, "aaa")
		val, ok := m.Get(key)
		if ok == false || val.(string) != "aaabbb" {
			t.Fatalf("%t %q - should be not false", ok, val)
		}
	}

	f(1)
	// f(1.0)
}

func TestMapS(t *testing.T) {
	t.Helper()
	f := func(key string) {
		m := NewMapS()
		m.Add(key, 1)
		m.Add(key, 1.0)
		m.Add(key, "aaabbb")
		if _, ok := m.Get(key); ok == true {
			t.Fatal("should be not ok")
		}
		m.Apply()
		m.Add(key, "aaa")
		val, ok := m.Get(key)
		if ok == false || val.(string) != "aaabbb" {
			t.Fatal("should be not false")
		}
	}
	f("a")
}
