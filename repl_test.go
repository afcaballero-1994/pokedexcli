package main

import (
	"testing"
)

type Case struct {
	input string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []Case{
		{
			input: "  hello world  ",
			expected: []string {"hello", "world"},
		},
		{
			input: "  Hello world  ",
			expected: []string {"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string {"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("lens does not match.\nExpected: %d\nGot: %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("\nExpected: %s\n Got: %s", expectedWord, word)
			}
		}
	}
}