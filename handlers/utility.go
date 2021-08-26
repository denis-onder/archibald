package handlers

import (
	"io"
	"net/http"
)

func Welcome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome!")
	})
}
