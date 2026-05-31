package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input	string
		expected	[]string
	}{
		{
			input: "   hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " Charmander Bulbasaur  PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: " 	abc dEf   IJk\n",
			expected: []string{"abc", "def", "ijk"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("length of the actual against the expected don't match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected %s, actual %s", expectedWord, word)
			}
		}
	}
}
