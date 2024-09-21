package main

import (
	"flag"
	"fmt"
	"log"

	"NHS_Number_Validation_/NHS_number_validation/nhshandling"
)

func main() {
	// define the command-line flags
	operation := flag.String("operation", "", "Operation - validate or generate")

	numPtr := flag.String("num", "", "NHS number to validate")

	// parse the flags
	flag.Parse()

	// validate the operation flag \
	if *operation != "validate" && *operation != "generate" {
		fmt.Println("invalid operation")
	}

	// initialize the service interface
	var nhsService nhshandling.Service
	// use the service struct that implements the above interface
	nhsService = nhshandling.NewService()

	// validate a given NHS number in case operation flag = validate
	if *operation == "validate" {
		isValid, err := nhsService.ValidateNHS(*numPtr)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The NHS number %s is valid - %v\n", *numPtr, isValid)
	}

	// generate new NHS number in case operation flag = generate
	if *operation == "generate" {
		fmt.Printf("Generated NHS Number: %s\n", nhsService.GenerateNHS())
	}
}
