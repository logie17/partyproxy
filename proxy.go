package main

import (
	"net/http"
	"html"
	"fmt"
)

func main() {
	var port = 3000

	fmt.Println("let the party begin at %d", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Woot! %q", html.EscapeString(r.URL.Path))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Ah shoot johny: %v", err)
	}
}
