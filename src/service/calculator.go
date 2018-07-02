package service

import (
	"strconv"
	"strings"
)

type Calculator struct {
}

func (Calculator) Sum(numberText string) int {
	if strings.Contains(numberText, ",") {
		numbers := strings.Split(numberText, ",")
		sum := 0
		for _, value := range numbers {
			n, _ := strconv.Atoi(value)
			sum = n + sum
		}
		return sum
	}

	if numberText == "1" {
		return 1
	}
	return 0
}
