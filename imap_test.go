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
func TestMapIterate(t *testing.T) {
	t.Helper()
	m := NewMap()
	for index, val := range []int{0, 1, 2, 3, 4, 5} {
		m.Add(val, index)
	}
	m.Apply()

	for obj := range m.Interate() {
		if obj.val.(int) != obj.key {
			t.Fatal("Iterate not working")
		}
	}
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

func TestMapIIterate(t *testing.T) {
	t.Helper()
	m := NewMapI()
	for index, val := range []int{0, 1, 2, 3, 4, 5} {
		m.Add(val, index)
	}
	m.Apply()

	for obj := range m.Interate() {
		if obj.val.(int) != obj.key {
			t.Fatal("Iterate not working")
		}
	}
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

func TestMapSIterate(t *testing.T) {
	t.Helper()
	m := NewMapS()
	for index, val := range []string{"a", "b", "c", "d", "e"} {
		m.Add(val, index)
	}
	m.Apply()

	for obj := range m.Interate() {
		switch obj.key {
		case "a":
			if obj.val.(int) != 0 {
				t.Fatal("Iterate not working")
			}
		case "b":
			if obj.val.(int) != 1 {
				t.Fatal("Iterate not working")
			}
		case "c":
			if obj.val.(int) != 2 {
				t.Fatal("Iterate not working")
			}
		case "d":
			if obj.val.(int) != 3 {
				t.Fatal("Iterate not working")
			}
		case "e":
			if obj.val.(int) != 4 {
				t.Fatal("Iterate not working")
			}
		}
		// t.Logf("%s - %d", obj.key, obj.val)
	}
}
