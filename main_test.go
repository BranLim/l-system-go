package main

import "testing"

func TestAB(t *testing.T) {
	const axiom = "A"
	const expected = "AB"
	var rules = BuildRules("A", "AB", "B", "A")

	var result = rules.Produce(axiom, 1)
	assert(expected, result, t)
}

func TestABA(t *testing.T) {
	const axiom = "A"
	const expected = "ABA"
	var rules = BuildRules("A", "AB", "B", "A")
	var actual = rules.Produce(axiom, 2)
	assert(expected, actual, t)
}

func TestABAAB(t *testing.T) {
	const axiom = "A"
	const expected = "ABAAB"
	var rules = BuildRules("A", "AB", "B", "A")
	var actual = rules.Produce(axiom, 3)
	assert(expected, actual, t)
}

func TestLevel7(t *testing.T) {
	const axiom = "A"
	const expected = "ABAABABAABAABABAABABAABAABABAABAAB"
	var rules = BuildRules("A", "AB", "B", "A")
	assert(expected, rules.Produce(axiom, 7), t)
}

func TestFractal(t *testing.T) {
	const axiom = "0"
	const expected = "1[0]0"
	var rules = BuildRules("1", "11", "0", "1[0]0")
	assert(expected, rules.Produce(axiom, 1), t)
	const expected2 = "1111[11[1[0]0]1[0]0]11[1[0]0]1[0]0"
	assert(expected2, rules.Produce(axiom, 3), t)

}

func assert(expected string, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
