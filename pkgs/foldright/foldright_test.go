package foldright

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoldRight(t *testing.T) {
	nums := []int{1, 2, 3}
	assert.Equal(t, []int{6, 6, 6}, FoldRight(nums, nums, nums))
}
