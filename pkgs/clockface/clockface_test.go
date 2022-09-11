package clockface

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		Time  time.Time
		Angle float64
	}{
		{Angle: 0.0, Time: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{Angle: math.Pi, Time: time.Date(2019, time.January, 1, 0, 0, 30, 0, time.UTC)},
		{Angle: math.Pi * (1.0 / 3), Time: time.Date(2019, time.January, 1, 0, 0, 10, 0, time.UTC)},
		{Angle: math.Pi * (1.0 / 2), Time: time.Date(2019, time.January, 1, 0, 0, 15, 0, time.UTC)},
		{Angle: math.Pi * (55.0 / 30), Time: time.Date(2019, time.January, 1, 0, 0, 55, 0, time.UTC)},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Seconds: %d / Angle: %.2f", c.Time.Second(), c.Angle), func(t *testing.T) {
			got := secondsInRadians(c.Time)
			assert.InDelta(t, c.Angle, got, 10e-16)
		})
	}
}
