package fact

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestMyFactorization_Work(t *testing.T) {
	tests := []struct {
		name   string
		input  Input
		output []string
	}{
		{
			name: "NormalTest",
			input: Input{
				NumsOfGoroutine: 3,
				Numbers:         []int{7, 13, 21},
			},
			output: []string{"7 = 7", "13 = 13", "21 = 3 * 7"},
		},
		{
			name: "NegativeTest",
			input: Input{
				NumsOfGoroutine: 2,
				Numbers:         []int{-1, -10, 10, 15, 20},
			},
			output: []string{
				"-1 = -1 * 1", "-10 = -1 * 2 * 5",
				"10 = 2 * 5", "15 = 3 * 5", "20 = 2 * 2 * 5"},
		},
		{
			name: "MoreGoroutineTest",
			input: Input{
				NumsOfGoroutine: 5,
				Numbers:         []int{7, 13, 21},
			},
			output: []string{"7 = 7", "13 = 13", "21 = 3 * 7"},
		},
		{
			name: "BigTest",
			input: Input{
				NumsOfGoroutine: 3,
				Numbers:         []int{7, 13, 11, 21, 99, 58, 90, 11},
			},
			output: []string{
				"7 = 7", "13 = 13", "21 = 3 * 7",
				"99 = 3 * 3 * 11", "58 = 2 * 29",
				"90 = 2 * 3 * 3 * 5", "11 = 11", "11 = 11"},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			buf := new(bytes.Buffer)
			factorization := NewFactorization()
			err := factorization.Work(tc.input, buf)
			require.NoError(t, err)
			for _, elem := range tc.output {
				require.True(t, strings.Contains(buf.String(), elem))
			}
		})
	}
}
