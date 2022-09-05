package utils

import (
	"github.com/sqweek/dialog"
)

func Loadfile() string {
	file, err := dialog.File().Filter("Text-Files", "txt").Title("Load Proxies").Load()
	if err != nil {
		panic(err)
	}
	return file
}
