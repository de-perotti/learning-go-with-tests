package tremetreme

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTremeTreme(t *testing.T) {
	tamtam := TremeTreme(3)
	expected := "bumbumbum"
	assert.Equal(t, tamtam, expected)
}

func BenchmarkTremeTreme(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		TremeTreme(3)
	}
}
