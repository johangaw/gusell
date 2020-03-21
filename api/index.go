package handler

import (
	"fmt"
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go on Now!</h1>")
}
