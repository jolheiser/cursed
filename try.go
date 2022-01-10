package cursed

// Try to do something
func Try(fn func()) (res tryResult) {
	defer func() {
		if err := recover(); err != nil {
			res = tryResult{err}
		}
	}()
	fn()
	return tryResult{}
}

// Catch any problems
func (tr tryResult) Catch(fn func(any)) catchResult {
	fn(tr.ex)
	return catchResult{} // Missed your chance to deal with the ex
}

// Finally do something before being done
func (tr tryResult) Finally(fn func()) {
	fn()
}

// Finally do something before being done
func (cr catchResult) Finally(fn func()) {
	fn()
}

// tryResult is the result of a Try
type tryResult struct {
	ex any
}

// catchResult is the result of a Catch
type catchResult struct{}
