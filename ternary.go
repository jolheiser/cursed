package cursed

// ᕈ is the ternary operator...sort of
func ᕈ[T any](expr bool, t, f T) T {
	if expr {
		return t
	}
	return f
}
