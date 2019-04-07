package main

import (
	"bytes"
	"fmt"
)

func solution(n string) (string, string) {
	var left bytes.Buffer
	var right bytes.Buffer
	for _, char := range n {
		if char == '4'{
			left.WriteRune('2')
			right.WriteRune('2')
		} else {
			left.WriteRune(char)
			if right.Len() > 0 {
				right.WriteRune('0')
			}
		}
	}
	return left.String(), right.String()
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n string
		fmt.Scan(&n)

		sum1, sum2 := solution(n)

		fmt.Printf("Case #%d: %s %s\n", testCase, sum1, sum2)
	}
}
