package main

import (
	"testing"
)

func TestTime(t *testing.T) {
	testCases := []struct {
		name string
		time string
		want string
	}{
		{
			name: "int",
			time: "1673349503212",
			want: "13:18:23",
		},
		{
			name: "iso",
			time: "2024-10-25T10:38:43.047213+03:00",
			want: "10:38:43",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := ts(tt.time)
			if tt.want != got {
				t.Errorf("error parsing timestamp: ts: %q want %q got %q", tt.time, tt.want, got)
			}
		})
	}
}
