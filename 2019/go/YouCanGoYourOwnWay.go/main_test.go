package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var solutionTests = []struct {
		input       string
		output  string
	}{
		{"EEEEE", "SSSSS"},
		{"EESSS", "SSEEE"},
		{"ESESE", "SESES"},
	}

	for _, tt := range solutionTests {
		actual := solution(tt.input)
		if actual != tt.output {
			t.Errorf("INPUT:%s\nGot:%s\n Want:%s", tt.input, actual, tt.output)
		}
	}
}
