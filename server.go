package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// type backendserver struct {
// 	url string
// 	islive bool
// }
// type ServerPool struct {
// 	backends []*Backend
// 	current  uint64
//   }
  
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URL:", r.URL.String())
	fmt.Println("Path:", r.URL.Path)
	query := r.URL.Query()
	fmt.Println("Query parameters:")
    for key, values := range query {
        for _, value := range values {
            fmt.Printf("  %s = %s\n", key, value)
        }
    }
	backendURL, _ := url.Parse("http://127.0.0.1:8000/")
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}