package cursed

import (
	"testing"

	"github.com/matryer/is"
)

func TestTernary(t *testing.T) {
	assert := is.New(t)

	a := 1

	b := ᕈ(a == 1, "foo", "bar")
	assert.Equal(b, "foo") // b should be foo

	c := ᕈ(a == 2, true, false)
	assert.Equal(c, false) // c should be false
}
