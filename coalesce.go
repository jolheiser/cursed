package cursed

// ᕈᕈ is the coalesce operator...sort of
func ᕈᕈ[T comparable](t1, t2 T) T {
	var z T
	return ᕈ(t1 == z, t2, t1)
}
