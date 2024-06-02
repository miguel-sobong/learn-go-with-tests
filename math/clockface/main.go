// run command:
// go run .\clockface\main.go > ./clockface/clock.svg

package main

import (
	clockface "learn-go-with-tests/math"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
