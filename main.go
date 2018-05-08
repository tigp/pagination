package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tigp/pagination/pagination"
)

func main() {
	args := os.Args[1:]
	argsList := []string{"current_page", "total_pages", "boundaries", "around"}
	var receivedArgs []int

	if len(args) < 4 {
		log.Fatalln("Not enough arguments provided. Please call with:", strings.Join(argsList, " "))
	}

	// parse arguments
	for i, a := range argsList {
		if args[i] == "" {
			log.Fatalf("%s argument is not provided", a)
		}

		res, err := strconv.Atoi(args[i])
		if err != nil {
			log.Fatalf("%s is not an integer", a)
		}

		receivedArgs = append(receivedArgs, res)
	}

	// create pagination
	res, err := pagination.Paginate(receivedArgs[0], receivedArgs[1], receivedArgs[2], receivedArgs[3])
	if err != nil {
		log.Fatalf("Coudn't create pagination: %v", err)
	}

	// print result
	fmt.Println(strings.Join(res, " "))
}
