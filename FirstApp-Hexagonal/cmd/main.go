package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FirstApp-Hexagonal/pkg/storage"

	"github.com/FirstApp-Hexagonal/pkg/rest"

	"github.com/FirstApp-Hexagonal/pkg/adding"
)

func main() {
	r, err := storage.SetupStorage()
	if err != nil {
		fmt.Println(err)
	}

	as := adding.NewService(r)

	router := rest.Handler(as)

	fmt.Println("Starting server on port 8080.....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
