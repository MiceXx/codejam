package main

import (
	"reflect"
	"testing"
)

func Test_isPangram(t *testing.T) {
	var solutionTests = []struct {
		input  string
		output bool
	}{
		{"CJQUIZKNOWBEVYOFDPFLUXALGORITHMS", true},
		{"SUBDERMATOGLYPHICFJKNQVWXZ", true},
		{"THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG", true},
		{"THEQUICKBROWNFOXJUMPSOVERTHELAYYDOG", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYY", false},
		{"HELLOWORLD", false},
	}

	for _, tt := range solutionTests {
		actual := isPangram(tt.input)
		if actual != tt.output {
			t.Errorf("INPUT:%s\nGot:%t\n Want:%t", tt.input, actual, tt.output)
		}
	}
}

func Test_eratosthenes(t *testing.T) {
	var solutionTests = []struct {
		input  int
		output []int
	}{
		{5, []int{2, 3, 5}},
		{11, []int{2, 3, 5, 7, 11}},
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range solutionTests {
		actual := eratosthenes(tt.input)
		if !reflect.DeepEqual(actual, tt.output) {
			t.Errorf("INPUT:%d\nGot:%v\n Want:%v", tt.input, actual, tt.output)
		}
	}
}
func Test_solution(t *testing.T) {
	var solutionTests = []struct {
		input1 int
		input2 []int
		output string
	}{
		{103,
			[]int{217, 1891, 4819, 2291, 2987, 3811, 1739, 2491, 4717, 445, 65, 1079, 8383, 5353, 901, 187, 649, 1003, 697, 3239, 7663, 291, 123, 779, 1007, 3551, 1943, 2117, 1679, 989, 3053},
			"CJQUIZKNOWBEVYOFDPFLUXALGORITHMS"},
		{10000,
			[]int{3292937, 175597, 18779, 50429, 375469, 1651121, 2102, 3722, 2376497, 611683, 489059, 2328901, 3150061, 829981, 421301, 76409, 38477, 291931, 730241, 959821, 1664197, 3057407, 4267589, 4729181, 5335543},
			"SUBDERMATOGLYPHICFJKNQVWXZ"},
		{100000000,
			[]int{3292937, 175597, 18779, 50429, 375469, 1651121, 2102, 3722, 2376497, 611683, 489059, 2328901, 3150061, 829981, 421301, 76409, 38477, 291931, 730241, 959821, 1664197, 3057407, 4267589, 4729181, 5335543},
			"SUBDERMATOGLYPHICFJKNQVWXZ"},
	}

	for _, tt := range solutionTests {
		actual := solution(tt.input1, tt.input2)
		if actual != tt.output {
			t.Errorf("\n--Got:%s\n Want:%s", actual, tt.output)
		}
	}
}
