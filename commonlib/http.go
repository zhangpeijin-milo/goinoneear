package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func hehehe(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hehehe")
}