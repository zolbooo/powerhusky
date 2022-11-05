package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Printf("Usage: %s <command> [arguments]\n", os.Args[0])
	fmt.Println("The commands are:")
	fmt.Println("\tinstall - install daemon")
	fmt.Println("\tuninstall - uninstall daemon and disable service")
}

func main() {
	// installCmd := flag.NewFlagSet("install", flag.ExitOnError)
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	default:
		printUsage()
		os.Exit(1)
	}
}
