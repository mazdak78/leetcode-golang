package main

import "strings"

func main() {
}

func romanToInt(s string) int {
	romansMap := make(map[string]int)
	romansMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	roman2DigitMap := make(map[string]int)
	roman2DigitMap = map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	if len(s) == 1 {
		return romansMap[s]
	}

	sum := 0
	slice := strings.Split(s, "")

	for {
		m := slice[:2]
		if roman2DigitMap[strings.Join(m, "")] != 0 {
			sum += roman2DigitMap[strings.Join(m, "")]
			slice = slice[2:]
		} else {
			sum += romansMap[strings.Join(slice[:1], "")]
			slice = slice[1:]
		}

		if len(slice) == 1 {
			sum += romansMap[strings.Join(slice[:1], "")]
			break
		}

		if len(slice) == 0 {
			break
		}
	}

	return sum
}
