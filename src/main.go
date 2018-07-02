package main

import (
	"fmt"
	"service"
)

func main() {
	cal := service.Calculator{}
	result := cal.Sum("9,8,9")
	fmt.Println(result)
}
