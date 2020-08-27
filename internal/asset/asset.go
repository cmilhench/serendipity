package asset

//go:generate go run gen.go

var store = make(map[string][]byte)

// Add asset data bytes
func Add(name string, content []byte) {
	store[name] = content
}

// Get returns asset bytes
func Get(name string) []byte {
	if f, ok := store[name]; ok {
		return f
	}
	return []byte{}
}

// Has checks if an asset exists in the store
func Has(name string) bool {
	if _, ok := store[name]; ok {
		return true
	}
	return false
}

// Keys returns a slice of keys in the store
func Keys() []string {
	keys := make([]string, 0, len(store))
	for k := range store {
		keys = append(keys, k)
	}
	return keys
}
