package main

import (
	"testing"
)

var TestCase = []struct {
	input  int
	output string
}{
	{
		input:  12,
		output: "XII",
	},
	{
		input:  10,
		output: "X",
	},
	{
		input:  79,
		output: "LXXIX",
	},
	{
		input:  225,
		output: "CCXXV",
	},
	{
		input:  845,
		output: "DCCCXLV",
	},
	{
		input:  2022,
		output: "MMXXII",
	},
	{
		input:  123,
		output: "CXXIII",
	},
	{
		input:  3999,
		output: "MMMCMXCIX",
	}, {
		input:  3090,
		output: "MMMXC",
	}, {
		input:  1,
		output: "I",
	}, {
		input:  23,
		output: "XXIII",
	}, {
		input:  100,
		output: "C",
	}, {
		input:  91,
		output: "XCI",
	}, {
		input:  1234,
		output: "MCCXXXIV",
	}, {
		input:  3211,
		output: "MMMCCXI",
	}, {
		input:  873,
		output: "DCCCLXXIII",
	}, {
		input:  1863,
		output: "MDCCCLXIII",
	}, {
		input:  1975,
		output: "MCMLXXV",
	}, {
		input:  1984,
		output: "MCMLXXXIV",
	}, {
		input:  1009,
		output: "MIX",
	}, {
		input:  1910,
		output: "MCMX",
	}, {
		input:  1099,
		output: "MXCIX",
	}, {
		input:  1555,
		output: "MDLV",
	}, {
		input:  1549,
		output: "MDXLIX",
	}, {
		input:  777,
		output: "DCCLXXVII",
	}, {
		input:  2222,
		output: "MMCCXXII",
	}, {
		input:  1111,
		output: "MCXI",
	},
	{
		input:  3303,
		output: "MMMCCCIII",
	},
	{
		input:  3350,
		output: "MMMCCCL",
	},
	{
		input:  444,
		output: "CDXLIV",
	},
	{
		input:  3777,
		output: "MMMDCCLXXVII",
	},
	{
		input:  3709,
		output: "MMMDCCIX",
	},
}

func Test_Driver(t *testing.T) {
	for _, testCase := range TestCase {
		output := intToRoman(testCase.input, 1)

		if output != testCase.output {
			t.Errorf("intToRoman output: %s != TestCase output %s, for input %+v", output, testCase.output, testCase.input)
		}
	}
}
