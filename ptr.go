package cursed

// NewPtr returns a new pointer to T, or an empty T if no args are provided
func NewPtr[T any](in ...T) *T {
	return ᕈ(len(in) > 0, func() *T { return &(&struct{ v T }{in[0]}).v }, func() *T { return &(&struct{ v T }{}).v })()
}
