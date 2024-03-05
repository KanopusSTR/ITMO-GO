package _defer

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestPrintSequence1(t *testing.T) {
	writer, _ := os.OpenFile("wut_output_1.tsv", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(writer *os.File) {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}(writer)

	PrintSequence1(writer)

	content, err := os.ReadFile("wut_output_1.tsv")
	if err != nil {
		panic(err)
	}

	require.Equal(t, "135642", string(content))
}

func TestPrintSequence2(t *testing.T) {
	writer, _ := os.OpenFile("wut_output_2.tsv", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(writer *os.File) {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}(writer)

	PrintSequence2(writer)

	content, err := os.ReadFile("wut_output_2.tsv")
	if err != nil {
		panic(err)
	}

	require.Equal(t, "1345287960", string(content))
}
