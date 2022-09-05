package utils

import (
	"fmt"
)

func Input(text string) string {
	var input string
	fmt.Print(text)
	fmt.Scanln(&input)
	return input
}
