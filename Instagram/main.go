package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kunalk27/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to the Instagram API!")
	r := router.Router()
	fmt.Println("The Server is getting started...")
	log.Fatal(http.ListenAndServe(":2700", r))
	fmt.Println("Listening at port 2700...")

}
