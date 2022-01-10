package cursed

import (
	"github.com/matryer/is"
	"testing"
)

func TestCoalesce(t *testing.T) {
	assert := is.New(t)

	type foo struct {
		Name string
	}

	b := ᕈᕈ("", "foo")
	assert.Equal(b, "foo") // b should be foo

	c := ᕈᕈ(true, false)
	assert.Equal(c, true) // c should be true

	dd := &foo{"test"}
	d := ᕈᕈ(nil, dd)
	assert.Equal(d, dd) // d should be &foo{"test"}

	e := ᕈᕈ(dd, nil)
	assert.Equal(e, dd) // e should be &foo{"test"}

	var nilf *foo
	f := ᕈᕈ(nil, nilf)
	assert.Equal(f, nilf) // f should be nilf
}
