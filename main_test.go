package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	if result != 3 {
		t.Errorf("Subtract(5, 2) = %d; want 3", result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	if result != 12 {
		t.Errorf("Multiply(3, 4) = %d; want 12", result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", result)
	}

	result = Divide(10, 0)
	if result != 0 {
		t.Errorf("Divide(10, 0) = %d; want 0", result)
	}
}

func TestToUpperCase(t *testing.T) {
	result := ToUpperCase("hello")
	if result != "HELLO" {
		t.Errorf("ToUpperCase('hello') = %s; want 'HELLO'", result)
	}
}
