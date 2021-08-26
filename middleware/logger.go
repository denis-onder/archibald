package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(h http.Handler) http.Handler {
	// Method, endpoint and timestamp
	logFunc := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n=> [%s] %s @%d\n", r.Method, r.URL.Path, time.Now().Unix())
	}

	return http.HandlerFunc(logFunc)
}
