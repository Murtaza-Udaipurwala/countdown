package main

import "flag"

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
