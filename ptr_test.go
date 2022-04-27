package cursed

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func TestNewPtr(t *testing.T) {
	assert := is.New(t)

	str := NewPtr[string]()
	assert.True(reflect.TypeOf(str).Kind() == reflect.Ptr)           // str should be a pointer
	assert.True(reflect.TypeOf(str).Elem().Kind() == reflect.String) // str should be a string

	i := NewPtr(5)
	assert.True(reflect.TypeOf(i).Kind() == reflect.Ptr)        // i should be a pointer
	assert.True(reflect.TypeOf(i).Elem().Kind() == reflect.Int) // i should be a string
	assert.True(*i == 5)                                        // *i should be 5
}
