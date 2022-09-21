package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"time"
)

type flushable interface {
	Flush() error
}

func flusher(w flushable, d time.Duration) {
	t := time.NewTicker(time.Second * 1)
	for range t.C {
		_ = w.Flush()
	}
}

func ts(t string) string {
	tt, err := time.Parse(time.RFC3339Nano, t)
	if err != nil {
		return "err"
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
	skip := map[string]bool{
		"logger":    true,
		"timestamp": true,
		"severity":  true,
	}

	reader := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(reader)

	// flush to std out every 5 seconds
	go flusher(w, time.Second*5)

	fp := color.New(color.FgWhite)
	mp := color.New(color.FgWhite, color.Bold)
	epk := color.New(color.FgCyan, color.Faint)
	epv := color.New(color.FgWhite, color.Faint)

	for scanner.Scan() {
		message := make(map[string]interface{})
		err := json.Unmarshal([]byte(scanner.Text()), &message)
		if err != nil {
			// if we can't parse just print it
			_, _ = fp.Println(scanner.Text())
			continue
		}

		if v, ok := message["severity"]; ok {
			lp := level(fmt.Sprintf("%s", v))
			_, _ = fp.Fprintf(w, "[")
			_, _ = lp.Fprintf(w, "%5s", v)
			_, _ = fp.Fprintf(w, "] ")
		}

		if v, ok := message["timestamp"]; ok {
			_, _ = fp.Fprintf(w, "%s ", ts(fmt.Sprintf("%s", v)))
		}

		if v, ok := message["logger"]; ok {
			_, _ = fp.Fprintf(w, "%-12s - ", v)
		}

		if v, ok := message["message"]; ok {
			_, _ = mp.Fprintf(w, "%s ", v)
		}

		for k, v := range message {
			if !skip[k] {
				_, _ = epk.Fprintf(w, "%s", k)
				_, _ = epv.Fprintf(w, "=")
				_, _ = epv.Fprintf(w, "%v ", v)
			}
		}

		_, _ = fp.Fprintf(w, "\n")
	}
}
