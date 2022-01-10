package cursed_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	. "github.com/jolheiser/cursed"

	"github.com/matryer/is"
)

func TestTry(t *testing.T) {
	assert := is.New(t)

	var ex any
	catcher := func(a any) {
		ex = a
	}

	Try(func() {
		panic("oops")
	}).Catch(catcher)
	assert.Equal(ex, "oops") // panic

	Try(func() {
		panic("oops")
	}).Catch(catcher).Finally(func() {
		ex = "finally"
	})
	assert.Equal(ex, "finally") // finally

	myErr := errors.New("uh oh")
	Try(func() {
		panic(myErr)
	}).Catch(catcher)
	assert.Equal(ex, myErr)

	Try(func() {
		var isNil *struct{ Field string }
		_ = isNil.Field
	}).Catch(catcher)
	assert.True(strings.Contains(fmt.Sprint(ex), "dereference")) // nil dereference
}
