package service

import "testing"

func Test_Sum_Input_Emply_Should_Be_Zero(t *testing.T) {
	cal := Calculator{}
	result := cal.Sum("")
	expected := 0
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}

func Test_Sum_Input_1_Should_Be_1(t *testing.T) {
	cal := Calculator{}
	result := cal.Sum("1")
	expected := 1
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}

func Test_Sum_Input_1_And_2_Should_Be_3(t *testing.T) {
	cal := Calculator{}
	result := cal.Sum("1,2")
	expected := 3
	if result != expected {
		t.Fatal("Expected ", expected, "but actual", result)
	}
}
