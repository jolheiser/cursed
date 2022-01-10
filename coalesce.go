package cursed

// ᕈᕈ is the coalesce operator...sort of
func ᕈᕈ[T comparable](t1, t2 T) T {
	var z T
	if t1 == z {
		return t2
	}
	return t1
}
