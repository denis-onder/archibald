package main

import (
	"fmt"
	"net/http"
	"os"
)

func serve() {
	port := ":" + os.Getenv("PORT")

	fmt.Printf("Archibald service running\nhttp://localhost%s/welcome\n", port)

	mux := http.NewServeMux()

	initializeRoutes(mux)

	http.ListenAndServe(port, mux)
}
