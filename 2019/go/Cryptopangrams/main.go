package main

import (
	"bytes"
	"fmt"
	"sort"
)

var alphas = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

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

//Implements sieve of eratosthenes which finds a slice of all prime numbers up to and including n
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

//Takes a map of primes and assigns alpha characters according to the order in sorted slice
func primesToMap(primesMap map[int]rune) map[int]rune {
	var tmp []int
	for k := range primesMap {
		tmp = append(tmp, k)
	}
	sort.Ints(tmp)
	for i, v := range tmp {
		primesMap[v] = alphas[i]
	}
	return primesMap
}

func solution(n int, ciphertext []int) string {
	var plaintext bytes.Buffer
	cipherPrimes := make([]int, len(ciphertext)+1)
	allPrimes := eratosthenes(n)
	selectedPrimes := make(map[int]rune)
	findDivisors := func(product int) (int, int) {
		for _, prime := range allPrimes {
			if product%prime == 0 {
				return prime, product / prime
			}
		}
		return 0, 0
	}

	for i, product := range ciphertext {
		var p1, p2 int
		if i > 0 {
			if product%cipherPrimes[i] == 0 {
				p1 = cipherPrimes[i]
				p2 = product / cipherPrimes[i]
			} else {
				cipherPrimes[i], cipherPrimes[i-1] = cipherPrimes[i-1], cipherPrimes[i]
			}
			p1 = cipherPrimes[i]
			p2 = product / cipherPrimes[i]
		} else {
			p1, p2 = findDivisors(product)
		}
		cipherPrimes[i] = p1
		cipherPrimes[i+1] = p2
		if _, ok := selectedPrimes[p1]; !ok {
			selectedPrimes[p1] = '.'
		}
		if _, ok := selectedPrimes[p2]; !ok {
			selectedPrimes[p2] = '.'
		}
	}
	selectedPrimes = primesToMap(selectedPrimes)
	for _, prime := range cipherPrimes {
		plaintext.WriteRune(selectedPrimes[prime])
	}
	return plaintext.String()
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n, l int
		fmt.Scan(&n, &l)
		ciphertext := make([]int, l)
		for i := range ciphertext {
			fmt.Scan(&ciphertext[i])
		}

		plaintext := solution(n, ciphertext)

		fmt.Printf("Case #%d: %s\n", testCase, plaintext)
	}
}
