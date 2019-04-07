package main

import (
	"bytes"
	"fmt"
)

var alphas = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func getNextPrime(n int) int {
	return 0
}

//isPangram checks if every character of the alphabet appears at least once in the string
//Only checks input string with ALL CAPS
func isPangram(s string) bool {
	var missing uint32 = (1 << 26) - 1
	for _, c := range s {
		var index uint32
		if 'A' <= c && c <= 'Z' {
			index = uint32(c - 'A')
		} else {
			continue
		}

		missing &^= 1 << index
		if missing == 0 {
			return true
		}
	}
	return false
}

func eratosthenes(n int) []int {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if integers[p] {
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

//primeFactors finds the letters corresponding to the factors of the given input or error if none exists
func primeFactorLetters(n int, primes []int) (rune, rune) {
	for i := 0; i < len(primes); i++ {
		for j := i + 1; j < len(primes); j++ {
			if (primes[i] * primes[j]) == n {
				return alphas[i], alphas[j]
			}
		}
	}
	return '.', '.'
}

//Given a set of primes, check if ciphertext can be decrypted and returns string or empty  string
func findPlaintext(primes []int, cipertext []int) string {
	var plaintext bytes.Buffer
	var checkValid func(idx int) bool
	checkValid = func(idx int) bool {
		if idx >= len(cipertext) {
			return true
		}
		letter1, letter2 := primeFactorLetters(cipertext[idx], primes)
		if letter1 != '.' && checkValid(idx+1) {
			plaintext.WriteRune(letter1)
			plaintext.WriteRune(letter2)
		}
		return false
	}
	_ = checkValid(0)
	return plaintext.String()
}

func solution(n int, cipertext []int) string {
	primes := eratosthenes(n)
	var start = 0
	var end = 25

	for i := 0; i < len(primes)-25; i++ {
		selectedPrimes := primes[start+i : end+i]
		plaintext := findPlaintext(selectedPrimes, cipertext)
		if plaintext != "" {
			return plaintext
		}
	}
	return ""
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n, l int
		fmt.Scan(&n, &l)
		cipertext := make([]int, l)
		for i := range cipertext {
			fmt.Scan(&cipertext[i])
		}

		plaintext := solution(n, cipertext)

		fmt.Printf("Case #%d: %s\n", testCase, plaintext)
	}
}
