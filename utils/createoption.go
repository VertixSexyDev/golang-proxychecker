package utils

import (
	"github.com/fatih/color"
)

func Createoption(n string, text string) {
	color.White("	[" + color.CyanString(n) + "] " + "- " + text)
}
