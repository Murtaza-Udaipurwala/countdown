package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type flags struct {
	seconds   uint
	minutes   uint
	hours     uint
	precision float64
	colorful  bool
}

func parseArgs() flags {
	seconds := flag.Uint("s", 0, "seconds")
	minutes := flag.Uint("m", 0, "minutes")
	hours := flag.Uint("h", 0, "hours")
	precision := flag.Float64("p", 1, "precision or how frequently output changes(in seconds)")
	colorful := flag.Bool("c", false, "colorful output")
	flag.Parse()

	return flags{
		seconds:   *seconds,
		minutes:   *minutes,
		hours:     *hours,
		precision: *precision,
		colorful:  *colorful,
	}
}

var colors = []color.Attribute{
	color.FgRed,
	color.FgGreen,
	color.FgYellow,
	color.FgBlue,
	color.FgMagenta,
	color.FgCyan,
	color.FgHiRed,
	color.FgHiGreen,
	color.FgHiYellow,
	color.FgHiBlue,
	color.FgHiMagenta,
	color.FgHiCyan,
}

func getColor() color.Attribute {
	return colors[rand.Intn(len(colors))]
}

func watch(colorful bool) {
	start := time.Now()
	for {
		delta := time.Now().Sub(start)
		seconds := int(delta.Seconds())
		minutes := int(delta.Minutes())
		hours := int(delta.Hours())

		if colorful {
			color.New(getColor()).Printf("\r%02d:%02d:%02d", hours, minutes, seconds)
		} else {
			fmt.Printf("\r%02d:%02d:%02d", hours, minutes, seconds)
		}

		time.Sleep(time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	flags := parseArgs()

	if flags.seconds == 0 && flags.minutes == 0 && flags.hours == 0 {
		watch(flags.colorful)
	}
}
