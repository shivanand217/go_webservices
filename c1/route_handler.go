package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {

	// get the request path of URL
	urlParams := mux.Vars(r)

	name := urlParams["user"]
	HelloMessage := "Hello " + name

	message := API{HelloMessage}

	// marshal the json value
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went Wrong !!")
	}

	fmt.Fprintf(w, string(output))
}

func main() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)

	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
