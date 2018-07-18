package imap

import (
	"sync"
)

//
// Map > map[interface{}]interface{}
//

type mapOfInterface map[interface{}]interface{}

func (m mapOfInterface) copy() *mapOfInterface {
	new := make(mapOfInterface)
	for key, value := range m {
		new[key] = value
	}
	return &new
}

// Map immutable struct map[interface{}]interface{}
type Map struct {
	mutex sync.Mutex
	imap  *mapOfInterface
	tmap  *mapOfInterface
}

// NewMap create immutable struct Map
func NewMap() *Map {
	im := Map{}
	im.imap = &mapOfInterface{}
	im.tmap = &mapOfInterface{}
	return &im
}

// Get value from the immutable map
func (m Map) Get(key interface{}) (val interface{}, ok bool) {
	val, ok = (*m.imap)[key]
	return val, ok
}

// Add a key inside the temporary map
func (m *Map) Add(key interface{}, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	(*m.tmap)[key] = value
}

// Delete a key inside the temporary map
func (m *Map) Delete(key interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(*m.tmap, key)
}

// Apply the temporary map into the immutable one
func (m *Map) Apply() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	tmp := m.tmap.copy()
	m.imap = m.tmap
	m.tmap = tmp
}

// Len calculate the len of immutable map
func (m Map) Len() int {
	return len((*m.imap))
}

//
// MapS > map[string]interface{}
//

// mapOfString define the type of map for string hash
type mapOfString map[string]interface{}

func (m mapOfString) copy() *mapOfString {
	new := make(mapOfString)
	for key, value := range m {
		new[key] = value
	}
	return &new
}

// MapS immutable struct for map[string]interface{}
type MapS struct {
	mutex sync.Mutex
	imap  *mapOfString
	tmap  *mapOfString
}

// NewMapS create immutable struct map[string]interface{}
func NewMapS() *MapS {
	im := MapS{}
	im.imap = &mapOfString{}
	im.tmap = &mapOfString{}
	return &im
}

// Get value from the immutable map
func (m MapS) Get(key string) (val interface{}, ok bool) {
	val, ok = (*m.imap)[key]
	return val, ok
}

// Add a key inside the temporary map
func (m *MapS) Add(key string, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	(*m.tmap)[key] = value
}

// Delete a key inside the temporary map
func (m *MapS) Delete(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(*m.tmap, key)
}

// Apply the temporary map into the immutable one
func (m *MapS) Apply() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	tmp := m.tmap.copy()
	m.imap = m.tmap
	m.tmap = tmp
}

// Len clculate the len of immutable map
func (m MapS) Len() int {
	return len((*m.imap))
}

//
// MapI > map[int]interface{}
//

// mapOfInt define the type of map for int hash
type mapOfInt map[int]interface{}

func (m mapOfInt) copy() *mapOfInt {
	new := make(mapOfInt)
	for key, value := range m {
		new[key] = value
	}
	return &new
}

// MapI immutable struct map[int]interface{}
type MapI struct {
	mutex sync.Mutex
	imap  *mapOfInt
	tmap  *mapOfInt
}

// NewMapI create immutable struct MapI
func NewMapI() *MapI {
	im := MapI{}
	im.imap = &mapOfInt{}
	im.tmap = &mapOfInt{}
	return &im
}

// Get value from the immutable map
func (m MapI) Get(key int) (val interface{}, ok bool) {
	val, ok = (*m.imap)[key]
	return val, ok
}

// Add a key inside the temporary map
func (m *MapI) Add(key int, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	(*m.tmap)[key] = value
}

// Delete a key inside the temporary map
func (m *MapI) Delete(key int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(*m.tmap, key)
}

// Apply the temporary map into the immutable one
func (m *MapI) Apply() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	tmp := m.tmap.copy()
	m.imap = m.tmap
	m.tmap = tmp
}

// Len clculate the len of immutable map
func (m MapI) Len() int {
	return len((*m.imap))
}

//
// AsyncMap > map[string]interface{} with mutex
//

type AsyncMap struct {
	mutex sync.Mutex
	m     map[string]interface{}
}

func NewAsyncMap() AsyncMap {
	return AsyncMap{
		m: make(map[string]interface{}),
	}
}

// Get value from the immutable map
func (m AsyncMap) Get(key string) (val interface{}) {
	m.mutex.Lock()
	val = m.m[key]
	m.mutex.Unlock()
	return val
}

// Add a key inside the temporary map
func (m AsyncMap) Add(key string, value interface{}) {
	m.mutex.Lock()
	m.m[key] = value
	m.mutex.Unlock()
}
