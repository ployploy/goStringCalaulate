package Test

import "testing"

func TestInputemplySholudOutputzero(t *testing.T) {
	cal := Calculate{}
	result := cal.sum("")
	expected := 0
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}
func TestInput1SholudOutput1(t *testing.T) {
	cal := Calculate{}
	result := cal.sum("1")
	expected := 1
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}

func TestInput1and2ShouldOutput3(t *testing.T) {
	cal := Calculate{}
	result := cal.sum("1,2")
	expected := 3
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}
