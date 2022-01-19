package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strings"
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

func coutdown(duration uint, sleep float64, colorful bool) {
	timeout := time.Duration(duration * 1e9)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	pFactor := (sleep / float64(duration)) * 100
	var i float64

Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			i += pFactor
			if i > 100 {
				i = 100
			}

			bar := strings.Repeat("=", int(i))

			if colorful {
				color.New(getColor()).Printf("\r[%-100s] %.2f%%", bar, i)
			} else {
				fmt.Printf("\r[%-100s] %.2f%%", bar, i)
			}

			time.Sleep(time.Duration(sleep * 1e9))
		}
	}

	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	flags := parseArgs()

	if flags.seconds == 0 && flags.minutes == 0 && flags.hours == 0 {
		watch(flags.colorful)
	} else {
		duration := flags.seconds + flags.minutes*60 + flags.hours*60*60
		coutdown(duration, flags.precision, flags.colorful)
	}
}
