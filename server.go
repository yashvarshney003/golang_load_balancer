package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
	"time"
	"sync"
)

type backendserver struct {
	url    string
	islive bool
	mux sync.RWMutex
}

func (b *backendserver) setlive(isLive bool) {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.islive = isLive
}

var server_list []backendserver = []backendserver{
	{url: "http://127.0.0.1:8000/", islive :true },
    {url: "http://127.0.0.1:8001/", islive: true},
	{url: "http://127.0.0.1:8002/", islive:true}}

var current uint64


func health_check(server backendserver) bool {
	resp, err := http.Get(server.url)
	if err != nil {
		fmt.Println("Server", server.url, "is down")
		return false
	}
	defer resp.Body.Close()
	fmt.Println("Server", server.url, "is up")
	return true

}
func regular_health_check(server_list []backendserver) {
	fmt.Println("Starting health check...")
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			fmt.Println("Starting health check...")
			for _, server := range server_list {
				if server.islive{
					if !health_check(server) {
						fmt.Println("Server", server.url, "is down")
						server.setlive(false)
					} else {
						fmt.Println("Server", server.url, "is up")
						server.setlive(true)
					}
			}
			fmt.Println("Health check completed")
		}
	}
}
}

func find_next_live_server() backendserver {
	current = get_server_index()
	len_of_server_list := len(server_list)
	for i := uint64(0); i < uint64(len_of_server_list); i++ {
		idx := (current + i) % uint64(len_of_server_list)
		if health_check(server_list[idx]) {
			server_list[idx].setlive(true)
			current =  idx
			return server_list[idx]
		}else{
			server_list[idx].setlive(true)
		}
	}
	return server_list[current]
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
	server_obj := find_next_live_server()
	fmt.Println("Current server:", server_obj)
	backendURL, _ := url.Parse(server_obj.url)
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	go regular_health_check(server_list)
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
