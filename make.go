package cursed

// Make makes a []T
func Make[T any]() []T {
	return make([]T, len([]T{}))
}
