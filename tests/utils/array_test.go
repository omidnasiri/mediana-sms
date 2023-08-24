package utils_test

import (
	"testing"

	"github.com/omidnasiri/mediana-sms/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestArray_Contains(t *testing.T) {
	type testCase struct {
		name  string
		array []string
		entry string
		want  bool
	}

	testCases := []testCase{
		{
			name:  "empty array",
			array: []string{},
			entry: "foo",
			want:  false,
		},
		{
			name:  "array contains entry string",
			array: []string{"foo", "bar"},
			entry: "foo",
			want:  true,
		},
		{
			name:  "array does not contain entry string",
			array: []string{"foo", "bar"},
			entry: "so",
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := utils.Contains(tc.array, tc.entry)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestArray_RemoveDuplicateValues(t *testing.T) {
	type testCase struct {
		name  string
		array []string
		want  []string
	}

	testCases := []testCase{
		{
			name:  "empty array",
			array: []string{},
			want:  []string{},
		},
		{
			name:  "array with no duplicated value",
			array: []string{"A", "B"},
			want:  []string{"A", "B"},
		},
		{
			name:  "array with duplicated value",
			array: []string{"A", "B", "A", "B", "C", "A", "", ""},
			want:  []string{"A", "B", "C", ""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := utils.RemoveDuplicateValues(tc.array)
			assert.Equal(t, tc.want, got)
		})
	}
}
