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
		{
			"AUDIT 1",
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			"AUDIT 2",
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			"AUDIT 3",
			"Don not be sad ,because sad backwards is das . And das not good",
			"Don not be sad, because sad backwards is das. And das not good",
		},
		{
			"AUDIT 4",
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
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
