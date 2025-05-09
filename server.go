package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type backendserver struct {
	url    string
	islive bool
}

// type ServerPool struct {
// 	backends []*Backend
// 	current  uint64
//   }

var server_list []backendserver = []backendserver{{"http://127.0.0.1:8000/", true}, {"http://127.0.0.1:8001/", true}, {"http://127.0.0.1:8002/", true}}
var current uint64

func health_check(server backendserver) bool {
	resp, err := http.Get(server.url)
	if err != nil {
		fmt.Println("Server", server.url, "is down")
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Server", server, "is up")
		return true
	}
	return false
}
func get_server_index() uint64 {
	return uint64(int(atomic.AddUint64(&current, uint64(1)) % uint64(len(server_list))))
}

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
	server_index := get_server_index()
	fmt.Println("Current server:", server_index)
	backendURL, _ := url.Parse(server_list[server_index].url)
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
