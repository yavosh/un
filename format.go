package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/fatih/color"
)

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
