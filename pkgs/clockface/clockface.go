package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SVGWriter(w io.Writer, tm time.Time) {
	const (
		svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
		bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
		svgEnd = `</svg>`
	)

	secondHandTag := func(p Point) string {
		return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
	}

	sh := SecondHand(tm)

	_, _ = io.WriteString(w, svgStart)
	_, _ = io.WriteString(w, bezel)
	_, _ = io.WriteString(w, secondHandTag(sh))
	_, _ = io.WriteString(w, svgEnd)
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
