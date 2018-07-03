package imap

// mapS define the type of map for string hash
type mapS map[string]interface{}

func (m mapS) copy() mapS {
	new := make(mapS)
	for key, value := range m {
		new[key] = value
	}
	return new
}

// mapI define the type of map for int hash
type mapI map[int]interface{}

func (m mapI) copy() mapI {
	new := make(mapI)
	for key, value := range m {
		new[key] = value
	}
	return new
}

// S immutable struct
type S struct {
	imap *mapS
	tmap mapS
}

// NewS immutable struct
func NewS() S {
	im := S{}
	im.imap = &mapS{}
	im.tmap = mapS{}
	return im
}

// Get value from the immutable map
func (i S) Get(key string) (val interface{}, ok bool) {
	val, ok = (*i.imap)[key]
	return val, ok
}

// Add a key inside the temporary map
func (i *S) Add(key string, value interface{}) {
	if len(i.tmap) == 0 {
		i.tmap = (*i.imap).copy()
	}
	i.tmap[key] = value
}

// Delete a key inside the temporary map
func (i *S) Delete(key string) {
	if len(i.tmap) == 0 {
		i.tmap = (*i.imap).copy()
	}
	delete(i.tmap, key)
}

// Apply the temporary map into the immutable one
func (i *S) Apply() {
	if len(i.tmap) != 0 {
		i.imap = &i.tmap
	} else {
		i.imap = nil
	}
}

// Len clculate the len of immutable map
func (i S) Len() int {
	return len((*i.imap))
}

// I immutable struct
type I struct {
	imap *mapI
	tmap mapI
}

// NewI immutable struct
func NewI() I {
	im := I{}
	im.imap = &mapI{}
	im.tmap = mapI{}
	return im
}

// Get value from the immutable map
func (i I) Get(key int) (val interface{}, ok bool) {
	val, ok = (*i.imap)[key]
	return val, ok
}

// Add a key inside the temporary map
func (i *I) Add(key int, value interface{}) {
	if len(i.tmap) == 0 {
		i.tmap = (*i.imap).copy()
	}
	i.tmap[key] = value
}

// Delete a key inside the temporary map
func (i *I) Delete(key int) {
	if len(i.tmap) == 0 {
		i.tmap = (*i.imap).copy()
	}
	delete(i.tmap, key)
}

// Apply the temporary map into the immutable one
func (i *I) Apply() {
	if len(i.tmap) != 0 {
		i.imap = &i.tmap
	} else {
		i.imap = nil
	}
}

// Len clculate the len of immutable map
func (i I) Len() int {
	return len((*i.imap))
}
