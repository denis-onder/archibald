package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

func logger(w http.ResponseWriter, r *http.Request) {
	// Method, endpoint and timestamp
	fmt.Printf("=> [%s] %s @%s\n", r.Method, r.URL.Host, time.Now().Unix())
}

func serve() {
	port := ":" + os.Getenv("PORT")

	fmt.Printf("Archibald service running\nhttp://localhost%s/welcome", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", logger)
	mux.HandleFunc("/welcome", welcome)

	http.ListenAndServe(port, mux)
}
