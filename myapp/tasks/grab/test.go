package grab

import (
	"fmt"
	"net/http"
//	"appengine"
//	"appengine/user"
)



func init() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}

	http.HandleFunc("/tasks/grab/", handler)


}
