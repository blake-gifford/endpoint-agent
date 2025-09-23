package main

import (
	"flag"
	"fmt"
	"log"
	"main/platform"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var token string
	var organization int

	flag.StringVar(&token, "token", "", "Auth Token")
	flag.IntVar(&organization, "organization", 0, "Organization ID")
	flag.Parse()

	os.Setenv("TOKEN", token)
	os.Setenv("ORGANIZATION", fmt.Sprintf("%d", organization))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Stopping osqueryd daemon...")
		platform.StopOsqueryDaemon()
		os.Exit(0)
	}()

	platform.Run()
	log.Printf("osqueryd daemon started successfully")

	select {}
}
