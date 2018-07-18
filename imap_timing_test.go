package imap

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkAsyncMap(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < b.N; i++ {
		m.Store("a", "a")
		if val, ok := m.Load("a"); ok != false && val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}

func BenchmarkAsyncMapRead(b *testing.B) {
	m := sync.Map{}
	m.Store("a", "a")
	for i := 0; i < b.N; i++ {
		if val, ok := m.Load("a"); ok != false && val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}

func BenchmarkAsyncMapWrite(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < b.N; i++ {
		m.Store("a", "a")
	}
	if val, ok := m.Load("a"); ok != false && val == nil {
		panic(fmt.Errorf("unexpected not found value"))
	}
}
func BenchmarkAsyncMapParallel(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Store("a", "a")
			if val, ok := m.Load("a"); ok != false && val == nil {
				panic(fmt.Errorf("unexpected not found value"))
			}
		}
	})
}

func BenchmarkMap(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.Add("a", "a")
		m.Apply()
		if val, ok := m.Get("a"); ok == false || val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}
func BenchmarkMapRead(b *testing.B) {
	m := NewMap()
	m.Add("a", "a")
	m.Apply()
	for i := 0; i < b.N; i++ {
		if val, ok := m.Get("a"); ok == false || val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}
func BenchmarkMapWrite(b *testing.B) {
	m := NewMap()
	for i := 0; i < b.N; i++ {
		m.Add("a", "a")
	}
	m.Apply()
	if val, ok := m.Get("a"); ok == false || val == nil {
		panic(fmt.Errorf("unexpected not found value"))
	}
}

func BenchmarkMapParallel(b *testing.B) {
	m := NewMap()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Add("a", "a")
			m.Apply()
			if val, ok := m.Get("a"); ok == false || val == nil {
				panic(fmt.Errorf("unexpected not found value"))
			}
		}
	})
}

func BenchmarkMapS(b *testing.B) {
	m := NewMapS()
	for i := 0; i < b.N; i++ {
		m.Add("a", "a")
		m.Apply()
		if val, ok := m.Get("a"); ok == false || val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}

func BenchmarkMapSRead(b *testing.B) {
	m := NewMapS()
	m.Add("a", "a")
	m.Apply()
	for i := 0; i < b.N; i++ {
		if val, ok := m.Get("a"); ok == false || val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}

func BenchmarkMapSWrite(b *testing.B) {
	m := NewMapS()
	for i := 0; i < b.N; i++ {
		m.Add("a", "a")
	}
	m.Apply()
}

func BenchmarkMapSParallel(b *testing.B) {
	m := NewMapS()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Add("a", "a")
			m.Apply()
			if val, ok := m.Get("a"); ok == false || val == nil {
				panic(fmt.Errorf("unexpected not found value"))
			}
		}
	})
}

func BenchmarkMapI(b *testing.B) {
	m := NewMapI()

	for i := 0; i < b.N; i++ {
		m.Add(1, "a")
		m.Apply()
		if val, ok := m.Get(1); ok == false || val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}

func BenchmarkMapIRead(b *testing.B) {
	m := NewMapI()
	m.Add(1, "a")
	m.Apply()
	for i := 0; i < b.N; i++ {
		if val, ok := m.Get(1); ok == false || val == nil {
			panic(fmt.Errorf("unexpected not found value"))
		}
	}
}

func BenchmarkMapIWrite(b *testing.B) {
	m := NewMapI()
	for i := 0; i < b.N; i++ {
		m.Add(1, "a")
	}
	m.Apply()
}

func BenchmarkMapIParallel(b *testing.B) {
	m := NewMapI()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Add(1, "a")
			m.Apply()
			if val, ok := m.Get(1); ok == false || val == nil {
				panic(fmt.Errorf("unexpected not found value"))
			}
		}
	})
}
