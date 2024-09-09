package main

import (
	"linktic-create-orders/cmd/api/application"
	"log"
	"net/http"
)

func main() {
	app, err := application.InitializeApplication()
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8081", app.Router)
	if err != nil {
		log.Fatal(err)
	}
}
