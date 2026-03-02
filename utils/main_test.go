package utils

import "testing"

func TestParseInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			"Punctuation and Up",
			"hello , world (up)",
			"hello, WORLD",
		},
		{
			"Multiple UP with number",
			"i am very (up, 2) happy",
			"i AM VERY happy",
		},
		{
			"Multiple UP with invalid number",
			"i am very (up, 2b) happy",
			"i am VERY happy",
		},
		{
			"Quotes and Punctuation",
			"it ' is  simple . '",
			"it 'is simple.'",
		},
		{
			"VOWEL",
			"There is no greater agony than bearing a untold story inside you. It is a hard feeling.",
			"There is no greater agony than bearing an untold story inside you. It is an hard feeling.",
		},
		{
			"PUNCTUATION",
			"Punctuation tests are ... kinda boring ,what do you think ? (up) ?",
			"Punctuation tests are... kinda boring, what do you THINK??",
		},
		{
			"PUNCTUATION + UP ",
			"hello ... (up)",
			"HELLO...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseInput(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
