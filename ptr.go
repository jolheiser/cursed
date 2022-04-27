package cursed

// NewPtr returns a new pointer to T, or an empty T if no args are provided
func NewPtr[T any](in ...T) *T {
	return map[bool]func() *T{true: func() *T { return &(&struct{ v T }{in[0]}).v }, false: func() *T { return &(&struct{ v T }{}).v }}[len(in) > 0]()
}
