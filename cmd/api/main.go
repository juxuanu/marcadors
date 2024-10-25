package main

import (
	"fmt"
	"marcadors/internal/db"
	"net/http"
	"os"
)

func main() {
	dbClient, err := db.BuildDbClient()
	if err != nil {
		panic(err)
	}

	for _, handler := range ApiHandlers {
		http.HandleFunc(handler.path, handler.function)
	}

	fmt.Println("Starting API server...")
	serverErr := http.ListenAndServe(":8080", nil)
	if serverErr != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error trying to start the server:\n%s", serverErr)
		return
	}
}
