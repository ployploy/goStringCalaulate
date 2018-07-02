package service

import "testing"

func TestInputemplySholudOutputzero(t *testing.T) {
	cal := Calculator{}
	result := cal.Sum("")
	expected := 0
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}

func TestInput1SholudOutput1(t *testing.T) {
	cal := Calculator{}
	result := cal.Sum("1")
	expected := 1
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}

func TestInput1and2ShouldOutput3(t *testing.T) {
	cal := Calculator{}
	result := cal.Sum("1,2")
	expected := 3
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}
