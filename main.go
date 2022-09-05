package main

import (
	"github.com/VertixSexyDev/goproxychecker/utils"
	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()
	utils.Logo()
	utils.Createoption("1", "Modules")
	var menuinput string = utils.Input("	>> ")
	if menuinput == "1" {
		screen.Clear()
		screen.MoveTopLeft()
		utils.Logo()
		utils.Createoption("1", "HTTP/HHTPS")
		utils.Createoption("2", "SOCKS4")
		utils.Createoption("3", "SOCKS5")
		var moduleinput string = utils.Input("	>> ")
		if moduleinput == "1" {
			screen.Clear()
			screen.MoveTopLeft()
			utils.Logo()
			path := utils.Loadfile()
			utils.FileReader(path, "1")
		} else if moduleinput == "2" {
			screen.Clear()
			screen.MoveTopLeft()
			utils.Logo()
			path := utils.Loadfile()
			utils.FileReader(path, "2")
		} else if moduleinput == "3" {
			screen.Clear()
			screen.MoveTopLeft()
			utils.Logo()
			path := utils.Loadfile()
			utils.FileReader(path, "3")
		}
	}
}
