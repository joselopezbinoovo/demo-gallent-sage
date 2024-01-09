package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func SetupServer() {
	router := mux.NewRouter()
	routerConfig(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Server running in port %s \n", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
