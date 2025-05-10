package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	server_name := os.Getenv("SERVER_NAME")
	if port == "" {
		port = "8000"
	}
	if server_name == "" {
		server_name = "localhost"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend on port %s\n and server_name %s", port,server_name)
	})
	fmt.Println("Backend server running on port", port)
	http.ListenAndServe(":" + port, nil)
}
