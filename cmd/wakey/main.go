package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

const (
	Version  = "0.0.1"
	APP_NAME = "wakey"
)

type Config struct {
	Duration time.Duration
	Interval time.Duration
}

func main() {
	// parsing flags
	cfg := parseFlags()

	width, height := robotgo.GetScreenSize()

	run(cfg, width, height)
}

func parseFlags() Config {
	var seconds int
	flag.IntVar(&seconds, "seconds", 0, "run duration in seconds")
	flag.IntVar(&seconds, "s", 0, "run duration in seconds (shorthand)")

	var minutes int
	flag.IntVar(&minutes, "minutes", 0, "run duration in minutes")
	flag.IntVar(&minutes, "m", 0, "run duration in minutes (shorthand)")

	var hours int
	flag.IntVar(&hours, "hours", 0, "run duration in hours")
	flag.IntVar(&hours, "h", 0, "run duration in hours (shorthand)")

	var interval int
	flag.IntVar(&interval, "interval", 5, "seconds between cursor movement")
	flag.IntVar(&interval, "i", 5, "seconds between cursor movement (shorthand)")

	var versionFlag bool
	flag.BoolVar(&versionFlag, "version", false, "show version")
	flag.BoolVar(&versionFlag, "v", false, "show version (shorthand)")

	flag.Parse()

	if versionFlag {
		fmt.Printf("%s %s\n", APP_NAME, Version)
		os.Exit(0)
	}

	duration := time.Duration(seconds)*time.Second + time.Duration(minutes)*time.Minute + time.Duration(hours)*time.Hour

	return Config{
		Duration: duration,
		Interval: time.Duration(interval) * time.Second,
	}
}

func run(cfg Config, width, height int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	start := time.Now()
	for {
		if shouldStop(cfg.Duration, start) {
			return
		}

		x, y := randomPosition(rng, width, height)

		robotgo.MoveSmooth(x, y)

		time.Sleep(cfg.Interval)
	}
}

func shouldStop(duration time.Duration, start time.Time) bool {
	if duration == 0 {
		return false
	}

	return time.Since(start) >= duration
}

func randomPosition(rng *rand.Rand, width, height int) (int, int) {
	const margin = 50
	x := rng.Intn(width-100) + margin
	y := rng.Intn(height-100) + margin

	return x, y
}
