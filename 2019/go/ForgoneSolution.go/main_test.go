package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var solutionTests = []struct {
		input       string
		outputLeft  string
		outputRight string
	}{
		{"4", "2", "2"},
		{"940", "920", "20"},
		{"4444", "2222", "2222"},
		{"14", "12", "2"},
		{"1440", "1220", "220"},
	}

	for _, tt := range solutionTests {
		actualLeft, actualRight := solution(tt.input)
		if actualLeft != tt.outputLeft || actualRight != tt.outputRight {
			t.Errorf("INPUT:%s\nGot:%s %s\n Want:%s %s", tt.input, actualLeft, actualRight, tt.outputLeft, tt.outputRight)
		}
	}
}
