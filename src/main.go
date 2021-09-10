package main

import (
	check "check-email-spoof/src/check"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	showName()
	for {
		showMenu()
		command := readOption()
		callFeature(command)
	}
}

func showName() {
	color.Blue("   ____ _               _      _____                 _ _   ____                     __ ")
	color.Blue("  / ___| |__   ___  ___| | __ | ____|_ __ ___   __ _(_) | / ___| _ __   ___   ___  / _|")
	color.Blue(" | |   | '_ \\ / _ \\/ __| |/ / |  _| | '_ ` _ \\ / _` | | | \\___ \\| '_ \\ / _ \\ / _ \\| |_ ")
	color.Blue(" | |___| | | |  __/ (__|   <  | |___| | | | | | (_| | | |  ___) | |_) | (_) | (_) |  _|")
	color.Blue("  \\____|_| |_|\\___|\\___|_|\\_\\ |_____|_| |_| |_|\\__,_|_|_| |____/| .__/ \\___/ \\___/|_|  ")
	color.Blue("                                                                |_|                    ")
}

func showMenu() {
	println("\n###### ENTER OPTION ######")
	println("[1] - Check vulnerability an host")
	println("[2] - Make send test email")
	println("[0] - Exit")
	println("##########################\n")
}

func callFeature(option int) {
	switch option {
	case 1:
		println("Enter with host:")
		host := readHost()
		check.CheckVulnerability(host)

	case 2:
		println("IN PROGRESS")
	case 0:
		println("Exit.")
		os.Exit(0)
	default:
		println("Command not found")
	}
}

func readHost() string {
	var host string
	fmt.Scan(&host)

	return host
}

func readOption() int {
	var command int
	fmt.Scan(&command)

	return command
}
