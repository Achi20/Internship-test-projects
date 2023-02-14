package main

import (
	"fmt"
	"go_test5/router"
	"log"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	err := r.Run()
	logFatal(err)
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
