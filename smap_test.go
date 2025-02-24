package smap

import (
	"strings"
	"testing"
)

func TestN(t *testing.T) {
	from := []int{1, 2, 3}
	to := N[int, float64](from)
	if len(to) != len(from) {
		t.Errorf("to length is not equal to from length")
	}
	for i, v := range to {
		if v != float64(from[i]) {
			t.Errorf("to[%d] is not equal to from[%d]", i, i)
		}
	}
}

func TestS(t *testing.T) {
	from := []string{"a", "b", "c"}
	to := S[string, []byte](from)
	if len(to) != len(from) {
		t.Errorf("to length is not equal to from length")
	}
	for i, v := range to {
		if v[0] != []byte(from[i])[0] {
			t.Errorf("to[%d] is not equal to from[%d]", i, i)
		}
	}
}

func TestSNil(t *testing.T) {
	from := [][]byte{[]byte("a"), nil, []byte("c")}
	to := S[[]byte, string](from)
	if len(to) != len(from) {
		t.Errorf("to length is not equal to from length")
	}
	for i, v := range to {
		if v != string(from[i]) {
			t.Errorf("to[%d] is not equal to from[%d]", i, i)
		}
	}
}

func TestF(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		f        func(string) string
		expected []string
	}{
		{
			name:     "uppercase transformation",
			input:    []string{"hello", "world"},
			f:        strings.ToUpper,
			expected: []string{"HELLO", "WORLD"},
		},
		{
			name:  "custom transformation",
			input: []string{"a", "b", "c"},
			f: func(s string) string {
				return s + s
			},
			expected: []string{"aa", "bb", "cc"},
		},
		{
			name:     "empty strings",
			input:    []string{"", "  ", "\t"},
			f:        strings.TrimSpace,
			expected: []string{"", "", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := F(tt.input, tt.f)
			if len(result) != len(tt.expected) {
				t.Errorf("got length %d, want %d", len(result), len(tt.expected))
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("index %d: got %q, want %q", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestFWithTypeConversion(t *testing.T) {
	input := []int{1, 2, 3}
	result := F(input, func(n int) string {
		return string(rune('A' + n - 1))
	})
	expected := []string{"A", "B", "C"}

	if len(result) != len(expected) {
		t.Errorf("got length %d, want %d", len(result), len(expected))
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("index %d: got %q, want %q", i, result[i], expected[i])
		}
	}
}
