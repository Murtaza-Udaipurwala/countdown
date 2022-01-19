package main

import (
	"flag"
	"fmt"
	"time"
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

func watch(colorful bool) {
	start := time.Now()
	for {
		delta := time.Now().Sub(start)
		seconds := int(delta.Seconds())
		minutes := int(delta.Minutes())
		hours := int(delta.Hours())

		fmt.Printf("\r%02d:%02d:%02d", hours, minutes, seconds)
		time.Sleep(time.Second)
	}
}

func main() {
	flags := parseArgs()

	if flags.seconds == 0 && flags.minutes == 0 && flags.hours == 0 {
		watch(flags.colorful)
	}
}
