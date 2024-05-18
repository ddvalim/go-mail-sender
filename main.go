package go_mail_sender

import (
	"fmt"
	"github.com/ddvalim/go-mail-sender/router"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", 5885), router.CreateRouter())
	if err != nil {
		log.Fatal(err)
	}
}
