package main

import (
	"net/http"
	"html"
	"fmt"
	"github.com/logie17/partyproxy/config"
)

func main() {
	var port = 3000
	
	fmt.Println("let the party begin at: ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Woot! %q", html.EscapeString(r.URL.Path))
	})

	configuration := config.LoadConfig()
	fmt.Printf("Configuration: %v\n", configuration)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Ah shoot johny: %v", err)
	}
}
