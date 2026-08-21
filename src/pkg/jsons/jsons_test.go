package jsons

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		Name string
		Case []byte
		Want any
	}{
		{
			"Object",
			[]byte("{\"test\": \"test\"}"),
			map[string]any{"test": "test"},
		},
		{
			"Char",
			[]byte("{\"test\": \"t\\\\nes\\\\tt\"}"),
			map[string]any{"test": "t\\nes\\tt"},
		},
		{
			"Array",
			[]byte("[{\"test\": \"test\"}]"),
			map[string]any{"test": "test"},
		},
		{
			"Invalid",
			[]byte("[{\"test\": }]"),
			nil,
		},
		{
			"Log",
			[]byte("test"),
			nil,
		},
		{
			"Objects",
			[]byte("{\"test\": \"test\"}" + "\n" + "{\"test\": \"test\"}"),
			nil,
		},
		{
			"Arrays",
			[]byte("{\"test\": \"test\"}" + "\n" + "[{\"test\": \"test\"}]"),
			nil,
		},
		{
			"Chars",
			[]byte("{\"test\": \"test\"}" + "\n" + "{\"test\": \"te\\\\tst\"}" + "\n" + "{\"test\": \"te\\\\nst\"}"),
			nil,
		},
		{
			"Invalids",
			[]byte("{\"test\": \"test\"}" + "\n" + "{\"test\": }" + "\n" + "{"),
			nil,
		},
		{
			"Logs",
			[]byte("test" + "\n" + "{\"test\": \"test\"}" + "\n" + "test"),
			map[string]any{"test": "test"},
		},
		{
			"Empty",
			[]byte("\n \n"),
			nil,
		},
		{
			"None",
			[]byte(""),
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var got any
			Unmarshal(test.Case, &got)

			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}
