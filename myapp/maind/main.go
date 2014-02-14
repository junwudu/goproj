package maind

import (
	"fmt"
	"net/http"
)

func init() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %v", "ys")
	}
	http.HandleFunc("/", handler)
}

