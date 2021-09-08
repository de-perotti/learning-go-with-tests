package adder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Adder(2, 2)
	expect := 4
	assert.Equal(t, expect, sum)
}
