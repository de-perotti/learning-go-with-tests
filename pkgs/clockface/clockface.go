package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
)

func secondsInRadians(tm time.Time) float64 {
	return math.Pi * (float64(tm.Second()) / 30)
}

func SecondHand(tm time.Time) Point {
	angle := secondsInRadians(tm)
	yPos := math.Cos(angle) * secondHandLength
	xPos := math.Sin(angle) * secondHandLength
	return Point{X: 150 + xPos, Y: 150 - yPos}
}
