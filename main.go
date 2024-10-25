package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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

var (
	fp  = color.New(color.FgWhite)
	mp  = color.New(color.FgWhite, color.Bold)
	epk = color.New(color.FgCyan, color.Faint)
	epv = color.New(color.FgWhite, color.Faint)

	skip = map[string]bool{
		"logger":    true,
		"level":     true,
		"timestamp": true,
		"time":      true,
		"severity":  true,
	}
)

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

func format(line string, w io.Writer) {
	message := make(map[string]interface{})
	err := json.Unmarshal([]byte(line), &message)

	if err != nil {
		_, _ = fp.Fprintf(w, line)
		return
	}

	if v, ok := message["severity"]; ok {
		lp := level(fmt.Sprintf("%s", v))
		_, _ = fp.Fprintf(w, "[")
		_, _ = lp.Fprintf(w, "%5s", v)
		_, _ = fp.Fprintf(w, "] ")
	}

	if v, ok := message["level"]; ok {
		lp := level(fmt.Sprintf("%s", v))
		_, _ = fp.Fprintf(w, "[")
		_, _ = lp.Fprintf(w, "%5s", v)
		_, _ = fp.Fprintf(w, "] ")
	}

	if v, ok := message["timestamp"]; ok {
		_, _ = fp.Fprintf(w, "%s ", ts(fmt.Sprintf("%s", v)))
	}

	if v, ok := message["time"]; ok {
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
