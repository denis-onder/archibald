package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

func serve() {
	port := ":" + os.Getenv("PORT")

	fmt.Printf("Archibald service running\nhttp://localhost%s/welcome", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", welcome)

	http.ListenAndServe(port, mux)
}
