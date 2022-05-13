// Authored by		: Markus Walker
// Date Modified	: 5/13/22

// Description		: Unit testing for the calculator.go program.

package main

import "testing"

// Create struct addTest for the TestAddition function, accompanied with some table-driven tests.
type addTest struct {
	num1, num2, result int
}

var addTests = []addTest{
	addTest{1, 1, 2},
	addTest{2, 10, 12},
	addTest{-15, -10, -25},
	addTest{-25, 5, -20},
}

// Create struct subTest for the TestSubtraction function, accompanied with some table-driven tests.
type subTest struct {
	num1, num2, result int
}

var subTests = []subTest{
	subTest{1, 1, 0},
	subTest{2, 10, -8},
	subTest{-15, -10, -5},
	subTest{-25, 5, -30},
}

// Create struct addTest for the TestMultiplication function, accompanied with some table-driven tests.
type multTest struct {
	num1, num2, result int
}

var multTests = []multTest{
	multTest{1, 1, 1},
	multTest{2, 10, 20},
	multTest{-15, -10, 150},
	multTest{-25, 5, -125},
}

// Create struct addTest for the TestDivision function, accompanied with some table-driven tests.
type divTest struct {
	num1, num2, result int
}

var divTests = []divTest{
	divTest{1, 1, 1},
	divTest{2, 10, 0},
	divTest{-15, -10, 1},
	divTest{-25, 5, -5},
}

func TestAddition(t *testing.T) {
	for _, test := range addTests {
		if output := Addition(test.num1, test.num2); output != test.result {
			t.Errorf("The output of %q DOES NOT equal the expected output of %q", output, test.result)
		}
	}
}

func TestSubtraction(t *testing.T) {
	for _, test := range subTests {
		if output := Subtraction(test.num1, test.num2); output != test.result {
			t.Errorf("The output of %q DOES NOT equal the expected output of %q", output, test.result)
		}
	}
}

func TestMultiplication(t *testing.T) {
	for _, test := range multTests {
		if output := Multiplication(test.num1, test.num2); output != test.result {
			t.Errorf("The output of %q DOES NOT equal the expected output of %q", output, test.result)
		}
	}
}

func TestDivision(t *testing.T) {
	for _, test := range divTests {
		if output := Division(test.num1, test.num2); output != test.result {
			t.Errorf("The output of %q DOES NOT equal the expected output of %q", output, test.result)
		}
	}
}
