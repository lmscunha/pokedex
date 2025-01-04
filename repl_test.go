package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   john paul ",
			expected: []string{"john", "paul"},
		},
		{
			input:    "   george   ringo  ",
			expected: []string{"george", "ringo"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length do not match - actual: %v expected: %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word don't match word: %s expectedWord: %s", word, expectedWord)
			}

		}
	}
}
