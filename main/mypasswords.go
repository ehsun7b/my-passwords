package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"

	"github.com/fatih/color"
)

const (
	// AppName is the short name for the app
	AppName = "MyPasswords Native"
	// VersionName is version name of MyPasswords app
	VersionName = "0.1.0"
	// VersionNumber is the integer version number of the app
	VersionNumber = 1

	versionSwitch      = "--version"
	versionSwithcShort = "-v"

	cmdLs   = "ls"
	cmdNew  = "new"
	cmdMenu = "menu"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		mainMenu()
	} else {
		argsMode(args)
	}
}

// if program is executed with command line switches
func argsMode(args []string) {
	if len(args) == 1 && (args[0] == versionSwitch ||
		args[0] == versionSwithcShort) {
		fmt.Printf("%s %s\n", AppName, VersionName)
	}
}

// main loop which listens to the user's inputs
func mainMenu() {
	handleSignit()

	exit := false

	reader := bufio.NewReader(os.Stdin)

	printMenu()
	for !exit {

		fmt.Print("$ ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")

		switch command {
		case cmdLs:
			doCmdLs()
		case cmdNew:
			doCmdLs()
		case cmdMenu:
			printMenu()
		default:
			doUnknownCmd(command)
		}
	}

	color.Unset()
}

func printMenu() {
	fmt.Println("Choose from the following options:")

	color.Set(color.FgYellow)
	fmt.Print(cmdLs)
	color.Unset()
	fmt.Println("\tlist the enteries")

	color.Set(color.FgYellow)
	fmt.Print(cmdNew)
	color.Unset()
	fmt.Println("\tnew entry")

	color.Set(color.FgYellow)
	fmt.Print("ls")
	color.Unset()
	fmt.Println("\tlist the enteries")

	color.Set(color.FgYellow)
	fmt.Print("menu")
	color.Unset()
	fmt.Println("\tlist of commands")

	color.Set(color.FgRed)
	fmt.Print("ctrl+c")
	color.Unset()
	fmt.Println("\texit")

	color.Unset()
}

func doCmdLs(cmd ...string) {
	fmt.Println("Do ls")
}

func doCmdNew(cmd ...string) {
	fmt.Println("Do new")
}

func doUnknownCmd(cmd ...string) {
	fmt.Printf("Un known command '%s'", cmd[0])
}

func handleSignit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("captured %v, cleaning up before exit", sig)

			// todo: cleanup before exit

			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()
}
