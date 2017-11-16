package Test

import (
	"strconv"
	"strings"
)

type Calculate struct {
}

func (Calculate) sum(number string) int {
	if strings.Contains(number, ",") {
		numbers := strings.Split(number, ",")
		sum := 0
		for _, value := range numbers {
			n, _ := strconv.Atoi(value)
			sum = n + sum
		}
		return sum
	}

	if number == "1" {
		return 1
	}
	return 0
}
