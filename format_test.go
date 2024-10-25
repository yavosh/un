package main

import (
	"bytes"
	"testing"
)

func TestFormat(t *testing.T) {

	testCases := []struct {
		name string
		line string
		want string
	}{
		{
			name: "with time",
			line: "{\"time\":\"2024-10-25T10:38:42.796872+03:00\",\"level\":\"INFO\",\"msg\":\"Shut down HTTP\"}",
			want: "[ INFO] 10:38:42 msg=Shut down HTTP \n",
		},
		{
			name: "with timestamp",
			line: "{\"timestamp\":\"2024-10-25T10:38:42.796872+03:00\",\"level\":\"INFO\",\"msg\":\"Shut down HTTP\"}",
			want: "[ INFO] 10:38:42 msg=Shut down HTTP \n",
		},
		{
			name: "with int",
			line: "{\"timestamp\":\"1673349503212\",\"level\":\"INFO\",\"msg\":\"Shut down HTTP\"}",
			want: "[ INFO] 13:18:23 msg=Shut down HTTP \n",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.Buffer{}
			format(tt.line, &b)
			got := b.String()

			if tt.want != got {
				t.Errorf("error formating: line: %q want %q got %q", tt.line, tt.want, got)
			}
		})
	}
}
