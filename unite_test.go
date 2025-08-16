package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUniteJson(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    [][]byte
		expected map[string]any
	}{
		{
			name: "single file",
			input: [][]byte{
				[]byte(`{"foo": 123, "bar": 234}`),
			},
			expected: map[string]any{
				"foo": 123.0,
				"bar": 234.0,
			},
		},
		{
			name: "override json",
			input: [][]byte{
				[]byte(`{"foo": 123, "bar": 234, "ss": ["x"], "dict1": {"a":"b"}, "dict2": {"c":"d"}}`),
				[]byte(`{"foo": 12345, "baz": "hoge", "ss": ["y"], "dict2": {"c":"e", "x":"y"}}`),
			},
			expected: map[string]any{
				"foo": 12345.0,
				"bar": 234.0,
				"baz": "hoge",
				"ss":  []any{"x", "y"},
				"dict1": map[string]any{
					"a": "b",
				},
				"dict2": map[string]any{
					"c": "e",
					"x": "y",
				},
			},
		},
		{
			name: "unite string to string slice",
			input: [][]byte{
				[]byte(`{"foo": "fuga"}`),
				[]byte(`{"foo": ["hoge", "piyo"]}`),
			},
			expected: map[string]any{
				"foo": []any{"hoge", "piyo"},
			},
		},
		{
			name: "unite empty json",
			input: [][]byte{
				[]byte(`{"foo": 123, "bar": 234}`),
				{},
			},
			expected: map[string]any{
				"foo": 123.0,
				"bar": 234.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := UniteJSON(tt.input)
			require.NoError(t, err)
			assert.EqualValues(t, tt.expected, got)
		})
	}
}

func TestUniteJson_Error(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    [][]byte
		expected map[string]any
	}{
		{
			name: "invalid json",
			input: [][]byte{
				[]byte(`{"foo": 123, "bar": 234,}`),
			},
			expected: map[string]any{
				"foo": 123.0,
				"bar": 234.0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := UniteJSON(tt.input)
			require.Error(t, err)
		})
	}
}
