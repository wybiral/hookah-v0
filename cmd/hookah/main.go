package main

import (
	"fmt"
	"github.com/wybiral/hookah-v0/internal/app"
	"os"
)

const version = "0.3.1"

func usage() {
	fmt.Println("NAME:")
	fmt.Println("   hookah - WebSocket pipeline tool\n")
	fmt.Println("USAGE:")
	fmt.Println("   hookah command [command options]\n")
	fmt.Println("VERSION:")
	fmt.Println("   " + version + "\n")
	fmt.Println("COMMANDS:")
	fmt.Println("   node        start new node")
	fmt.Println("   recv        pipe from node to local stdout")
	fmt.Println("   send        pipe local stdout to node")
	fmt.Println("   version     print current version")
	fmt.Println("")
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
		return
	}
	switch args[0] {
	case "node":
		app.NodeMain(args[1:])
	case "recv":
		app.RecvMain(args[1:])
	case "send":
		app.SendMain(args[1:])
	case "version":
		fmt.Println(version)
	default:
		usage()
		return
	}
}
