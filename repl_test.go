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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello                          world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  world worl  ",
			expected: []string{"hello", "world", "worl"},
		},
		{
			input:    "  helloworld  ",
			expected: []string{"helloworld"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("test failed %v", actual)
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("test failed")
				t.FailNow()
			}
		}
	}
}
