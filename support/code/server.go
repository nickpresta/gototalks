package main

import (
	"fmt"
	"net/http"
)

// START OMIT
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 世界")
	})
	http.ListenAndServe(":4000", nil)
}
// END OMIT
