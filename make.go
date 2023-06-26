package cursed

import "unsafe"

// Make makes a []T, if an integer is passed in then it make the slice n big
func Make[T any](n ...int) []T {
	var cap int
	switch len(n) {
	case 0:
		cap = len([]T{})
	case 1:
		cap = **(**int)(unsafe.Pointer(&n))
	default:
		cap = len(n)
	}
	return append([]T(nil), make([]T, cap)...)
}
