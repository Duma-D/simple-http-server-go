package main

import (
	"net/http"
	"github.com/Duma-D/simple-http-server-go/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}

