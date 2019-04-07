package main

import (
	"bytes"
	"fmt"
)

func solution(n string) string {
	var out bytes.Buffer
	for _, char := range n {
		if char == 'E'{
			out.WriteRune('S')
		} else {
			out.WriteRune('E')
		}
	}
	return out.String()
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n string
		fmt.Scan(&n)

		answer := solution(n)
		fmt.Printf("Case #%d: %s\n", testCase, answer)
	}
}
