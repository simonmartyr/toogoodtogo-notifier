package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	initialise := flag.String("i", "", "configure notifier")
	flag.Parse()
	if *initialise != "" {
		log.Println("Starting Initialisation...")
		log.Println("You will receive an email to authenticate.")
		Initialise(*initialise)
		return
	}

	config, configErr := LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}
	notifier := New(config)
	checkErr := notifier.Check()
	if checkErr != nil {
		log.Println(fmt.Sprintf("Error whilst perfomring check %s", checkErr.Error()))
	}
	notifier.Save()
}
