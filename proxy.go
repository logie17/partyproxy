package main

import (
	"net/http"
	"html"
	"fmt"
	"encoding/json"
	"os"
)

type MapType struct {
	Incoming_Url string
	Url_Pool []string
}

type Configuration struct {
	Maps []MapType
}

func main() {
	var port = 3000

	fmt.Println("let the party begin at: ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"Woot! %q", html.EscapeString(r.URL.Path))
	})

	file, _ := os.Open("./conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	fmt.Printf("Configuration: %v\n", configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Println("Ah shoot johny: %v", err)
	}
}
