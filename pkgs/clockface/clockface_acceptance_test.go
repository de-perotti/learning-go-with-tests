package clockface_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"

	"github.com/de-perotti/learning-go-with-tests/pkgs/clockface"
)

func TestSecondHandAtTimes(t *testing.T) {
	times := []struct {
		time  time.Time
		point clockface.Point
	}{
		{
			time:  time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			point: clockface.Point{X: 150, Y: 60},
		},
		{
			time:  time.Date(2019, time.January, 1, 0, 0, 30, 0, time.UTC),
			point: clockface.Point{X: 150, Y: 240},
		},
		{
			time:  time.Date(2019, time.January, 1, 0, 0, 15, 0, time.UTC),
			point: clockface.Point{X: 240, Y: 150},
		},
		{
			time:  time.Date(2019, time.January, 1, 0, 0, 45, 0, time.UTC),
			point: clockface.Point{X: 60, Y: 150},
		},
	}
	for _, test := range times {
		t.Run(test.time.String(), func(t *testing.T) {
			got := clockface.SecondHand(test.time)
			assert.InDelta(t, test.point.Y, got.Y, 10e-10)
			assert.InDelta(t, test.point.X, got.X, 10e-10)
		})
	}
}

func TestSecondHandTag(t *testing.T) {
	tm := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)

	var b strings.Builder
	clockface.SVGWriter(&b, tm)
	got := b.String()
	want := `<line x1="150" y1="150" x2="150.000" y2="60.000"`

	assert.Contains(t, got, want)
}
