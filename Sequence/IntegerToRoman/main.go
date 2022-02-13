package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9.
    X can be placed before L (50) and C (100) to make 40 and 90.
    C can be placed before D (500) and M (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral.



Example 1:

Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.

Example 2:

Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

Example 3:

Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.



Constraints:

    1 <= num <= 3999


*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Input Decimal Integer: ")
		scanner.Scan()

		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		if input < 1 || input > 4000 {
			fmt.Println("Inputted integer must be between 1 <= x <= 3999")
			continue
		}

		fmt.Printf("Decimal Input: %s\n", scanner.Text())
		fmt.Printf("Roman Numeral Input: %s\n", intToRoman(input, 1))
	}
}

func intToRoman(remaining, power int) string {
	decimal := remaining % 10

	if decimal > 0 {
		return intToRoman(remaining/10, power+1) + computeRomanSymbol(power, decimal)
	}

	if decimal == 0 && power == 1 {
		return RomanSymbolMap[remaining]
	}

	return ""
}

func romanToInt(input string) int { return 0 }

func computeRomanSymbol(power, number int) string {
	high := int(math.Pow10(power))
	midpoint := high / 2
	low := int(math.Pow10(power - 1))

	switch scaledNumber := number * (int(math.Pow10(power - 1))); scaledNumber {
	case 4, 40, 400:
		return RomanSymbolMap[low] + RomanSymbolMap[midpoint]
	case 9, 90, 900:
		return RomanSymbolMap[low] + RomanSymbolMap[high]
	default:
		if symbol, ok := RomanSymbolMap[scaledNumber]; ok {
			return symbol
		}
		// either 3, 2, or 1
		if scaledNumber > midpoint {
			iterAmount := (scaledNumber - midpoint) / int(math.Pow10(power-1))
			// TODO: bug here, need to scale the larger symbol, 123 -> CXXIII, but this will compute CXIII
			return RomanSymbolMap[midpoint] + scale(RomanSymbolMap[low], iterAmount)
		}

		if scaledNumber < midpoint {
			iterAmount := (scaledNumber - low + 1) / int(math.Pow10(power-1))
			return scale(RomanSymbolMap[low], iterAmount)
		}
		// shouldn't be reached
		return ""
	}
}

func scale(symbol string, amount int) string {
	if amount > 0 {
		return scale(symbol, amount-1) + symbol
	}

	return ""
}

var RomanSymbolMap = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
