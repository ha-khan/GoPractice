package main

import (
	"testing"
)

var TestCase = []struct {
	input  int
	output string
}{
	{
		input:  2,
		output: "0102",
	},
	{
		input:  1,
		output: "01",
	}, {
		input:  5,
		output: "0102030405",
	}, {
		input:  10,
		output: "010203040506070809010",
	},
}

func Test_Driver(t *testing.T) {
	for _, testCase := range TestCase {
		output := driver(testCase.input)

		if output != testCase.output {
			t.Errorf("intToRoman output: %s != TestCase output %s, for input %+v", output, testCase.output, testCase.input)
		}
	}
}
