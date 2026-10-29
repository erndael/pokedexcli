package main

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"simple":               {input: "hello world", want: []string{"hello", "world"}},
		"empty string":         {input: "", want: []string{}},
		"leading whitespaces":  {input: "      hello world", want: []string{"hello", "world"}},
		"trailing whitespaces": {input: "hello world      ", want: []string{"hello", "world"}},
		"no whitespace":        {input: "helloworld", want: []string{"helloworld"}},
		"only whitespaces":     {input: "    ", want: []string{}},
		"multiple whitespaces": {input: "hello      world", want: []string{"hello", "world"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			if !slices.Equal(tc.want, got) {
				t.Fatalf("expected: %q, got %q", tc.want, got)
			}
		})
	}
}
