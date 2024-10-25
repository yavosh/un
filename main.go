package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type flushable interface {
	Flush() error
}

func flusher(w flushable, d time.Duration) {
	t := time.NewTicker(d)
	for range t.C {
		_ = w.Flush()
	}
}

func ts(t string) string {
	if ux, err := strconv.ParseInt(t, 10, 64); err == nil {
		return time.UnixMilli(ux).Format("15:04:05")
	}

	tt, err := time.Parse(time.RFC3339Nano, t)
	if err != nil {
		return "ts err"
	}

	return tt.Format("15:04:05")
}

func level(l string) *color.Color {

	switch strings.ToLower(l) {
	case "debug", "trace":
		return color.New(color.FgWhite)
	case "info":
		return color.New(color.FgGreen, color.Bold)
	case "warn":
		return color.New(color.FgYellow, color.Bold)
	case "error", "err", "fatal":
		return color.New(color.FgRed, color.Bold)
	default:
		return color.New(color.FgWhite, color.Faint)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(reader)

	// flush to std out every 5 seconds
	go flusher(w, time.Second*5)

	for scanner.Scan() {
		format(scanner.Text(), w)

	}
}
