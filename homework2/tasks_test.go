package main

import (
	"fmt"
	"slices"
	"testing"
)

var fibonacciTests = []struct {
	inputValue int
	expected   int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{19, 4181},
}

func TestFibonacciIterative(t *testing.T) {
	// The execution loop
	for _, testCase := range fibonacciTests {
		t.Run(fmt.Sprintf("F(%d)=%d.", testCase.inputValue, testCase.expected), func(t *testing.T) {
			ans := FibonacciIterative(testCase.inputValue)
			if ans != testCase.expected {
				t.Error(fmt.Sprintf("Got %d, expected %d", ans, testCase.expected))
			}
		})
	}
}

func TestFibonacciRecursive(t *testing.T) {
	// The execution loop
	for _, testCase := range fibonacciTests {
		t.Run(fmt.Sprintf("F(%d)=%d.", testCase.inputValue, testCase.expected), func(t *testing.T) {
			ans := FibonacciRecursive(testCase.inputValue)
			if ans != testCase.expected {
				t.Error(fmt.Sprintf("Got %d, expected %d", ans, testCase.expected))
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	primeNumbers := []int{
		2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
	}

	for i := range 50 {
		isPrime := slices.Contains(primeNumbers, i)
		t.Run(fmt.Sprintf("F(%d)=%t", i, isPrime), func(t *testing.T) {
			ans := IsPrime(i)
			if ans != isPrime {
				t.Error(fmt.Sprintf("For %d got %t, expected %t", i, ans, isPrime))
			}
		})
	}
}

func TestIsBinaryPalindrome(t *testing.T) {
	testCases := []struct {
		inputValue int
		expected   bool
	}{
		{7, true},
		{5, true},
		{6, false},
		{8, false},
		{9, true},
		{10, false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("F(%d)=%t.", testCase.inputValue, testCase.expected), func(t *testing.T) {
			ans := IsBinaryPalindrome(testCase.inputValue)
			if ans != testCase.expected {
				t.Error(fmt.Sprintf("Got %t, expected %t", ans, testCase.expected))
			}
		})
	}
}

func TestValidParentheses(t *testing.T) {
	testCases := []struct {
		inputValue string
		expected   bool
	}{
		{"[{}]", true},
		{"[{]}", false},
		{"[[[]]]", true},
		{"[]]]", false},
		{"]", false},
		{"", true},
		{"[", false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("F('%s')=%t.", testCase.inputValue, testCase.expected), func(t *testing.T) {
			ans := ValidParentheses(testCase.inputValue)
			if ans != testCase.expected {
				t.Error(fmt.Sprintf("Got %t, expected %t", ans, testCase.expected))
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	testCases := []struct {
		inputValue string
		expected   int
	}{
		{"0", 0},
		{"1", 1},
		{"01", 1},
		{"00001", 1},
		{"00101", 5},
		{"10101", 21},
		{"101010", 42},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("F('%s')=%d.", testCase.inputValue, testCase.expected), func(t *testing.T) {
			ans := Increment(testCase.inputValue)
			if ans != testCase.expected {
				t.Error(fmt.Sprintf("Got %d, expected %d", ans, testCase.expected))
			}
		})
	}
}
