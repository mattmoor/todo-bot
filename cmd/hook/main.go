package main

import (
	"net/http"

	"github.com/mattmoor/todo-bot/pkg/github"
)

func main() {
	http.HandleFunc("/", github.Handler)
	http.ListenAndServe(":8080", nil)
}
