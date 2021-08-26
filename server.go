package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func welcome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome!")
	})
}

func logger(h http.Handler) http.Handler {
	// Method, endpoint and timestamp
	logFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n=> [%s] %s @%d\n", r.Method, r.URL.Path, time.Now().Unix())
	}

	return http.HandlerFunc(logFunc)
}

func serve() {
	port := ":" + os.Getenv("PORT")

	fmt.Printf("Archibald service running\nhttp://localhost%s/welcome", port)

	mux := http.NewServeMux()
	mux.Handle("/welcome", logger(welcome()))

	http.ListenAndServe(port, mux)
}
