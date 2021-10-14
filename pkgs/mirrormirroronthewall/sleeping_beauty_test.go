package mirrormirroronthewall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalkToBall(t *testing.T) {
	printedStrings := []string{}
	mockFn := func(s string) {
		printedStrings = append(printedStrings, s)
	}

	princess := map[string]interface{}{
		"names": []string{"rolezao", "bailao"},
		"role":  2,
		"bla": map[string]interface{}{
			"bli": "foo",
			"bar": []string{"foo", "baz"},
		},
	}

	expected := []string{
		"rolezao",
		"bailao",
		"foo",
		"foo",
		"baz",
	}

	walk(princess, mockFn)

	assert.Equal(t, expected, printedStrings)
}
