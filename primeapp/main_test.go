package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Could not create a pipe: %v", err.Error())
	}
	os.Stdout = w
	prompt()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)

	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected -> but got %s", string(out))
	}
}
func Test_Intro(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf(err.Error())
	}

	os.Stdout = w

	intro()
	_ = w.Close()

	out, _ := io.ReadAll(r)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if string(out) != "Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> " {
		t.Errorf("incorrect output: expected Is it Prime?\n------------\nEnter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n-> but got '%s'", string(out))
	}
}
func Test_checkNumbers(t *testing.T) {
	testCases := []struct {
		input        string
		expected     string
		expectedDone bool
	}{
		{"2\n", "2 is a prime number!", false},
		{"3\n", "3 is a prime number!", false},
		{"4\n", "4 is not a prime number because it is divisible by 2!", false},
		{"-1\n", "Negative numbers are not prime, by definition!", false},
		{"abc\n", "Please enter a whole number!", false},
		{"q\n", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tc.input))
			result, done := checkNumbers(scanner)

			if result != tc.expected {
				t.Errorf("checkNumbers(%s) = %s; expected %s", tc.input, result, tc.expected)
			}

			if done != tc.expectedDone {
				t.Errorf("checkNumbers(%s) = %t; expected %t", tc.input, done, tc.expectedDone)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
