package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestIsNumInString(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "a4bc2d5e", expected: false},
		{input: "abc13cd", expected: true},
		{input: "", expected: false},
		{input: "aaa0b", expected: false},
		{input: "d\n5abc", expected: false},
		{input: "3ab10c", expected: true},
		{input: "45", expected: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := isNumInString(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestStringStartWithDigit(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{input: "a4bc2d5e", expected: false},
		{input: "abc13cd", expected: false},
		{input: "", expected: false},
		{input: "aaa0b", expected: false},
		{input: "d\n5abc", expected: false},
		{input: "3ab10c", expected: true},
		{input: "45", expected: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := stringStartWithDigit(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
