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
	seconds := flag.Int("seconds", 0, "run duration in seconds")
	minutes := flag.Int("minutes", 0, "run duration in minutes")
	hours := flag.Int("hours", 0, "run duration in hours")
	interval := flag.Int("interval", 5, "seconds between cursor movement")
	versionFlag := flag.Bool("version", false, "show version")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s %s\n", APP_NAME, Version)
		os.Exit(0)
	}

	duration := time.Duration(*seconds)*time.Second + time.Duration(*minutes)*time.Minute + time.Duration(*hours)*time.Hour

	return Config{
		Duration: duration,
		Interval: time.Duration(*interval) * time.Second,
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
