package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kardianos/service"
	"github.com/zolbooo/powerhusky/daemon"
)

func printUsage() {
	fmt.Printf("Usage: %s <command> [arguments]\n", os.Args[0])
	fmt.Println("The commands are:")
	fmt.Println("\tinstall - install daemon")
	fmt.Println("\tuninstall - uninstall daemon and disable service")
}

func runService(svc service.Service, daemonSvc *daemon.Service) {
	logger, err := svc.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	daemonSvc.Logger = logger
	err = svc.Run()
	if err != nil {
		logger.Error(err)
		return
	}
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	daemonSvc := &daemon.Service{}
	svc, err := service.New(daemonSvc, daemon.ServiceConfig)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "install":
		err = svc.Install()
		if err != nil {
			fmt.Printf("Failed to install service: %v\n", err)
			os.Exit(2)
		}
		os.Exit(0)
	case "uninstall":
		err = svc.Uninstall()
		if err != nil {
			fmt.Printf("Failed to uninstall service: %v\n", err)
			os.Exit(2)
		}
		os.Exit(0)
	default:
		if service.Interactive() {
			printUsage()
			os.Exit(1)
		}
	}

	runService(svc, daemonSvc)
}
