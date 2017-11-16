package Test

import "testing"

func TestInputemplySholudOutputzero(t *testing.T) {
	cal := Calculate{}
	result := cal.add("")
	expected := "0"
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}
