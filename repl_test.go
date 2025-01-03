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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello,world",
			expected: []string{"hello,world"},
		},
		{
			input:    "  HELLO WORLD  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		assertLength(t, len(actual), len(c.expected), c.expected)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			assertWord(t, word, expectedWord)
		}
	}
}

func assertLength(t testing.TB, actualLength, expectedLength int, expected []string) {
	t.Helper()
	if actualLength != expectedLength {
		t.Errorf("got length %d but wanted length %d when given %v", actualLength, expectedLength, expected)
	}
}

func assertWord(t testing.TB, actualWord, expectedWord string) {
	if actualWord != expectedWord {
		t.Errorf("got %q but was expecting %q", actualWord, expectedWord)
	}
}
