package main

import (
	"fmt"
	"github.com/ddvalim/go-mail-sender/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println(fmt.Sprintf("Server running on %d", 5885))

	err := http.ListenAndServe(fmt.Sprintf(":%d", 5885), router.CreateRouter())
	if err != nil {
		log.Fatal(err)
	}
}
