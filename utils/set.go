package utils

// Set implements the Set data structure where keys can be anything that map
type Set struct {
	cache map[interface{}]interface{}
}

// Put inserts the given key into set and overwrites if already existing
func (s *Set) Put(key interface{}) {
	s.cache[key] = struct{}{}
}
