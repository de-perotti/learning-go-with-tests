package main

import (
	"github.com/de-perotti/learning-go-with-tests/pkgs/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
