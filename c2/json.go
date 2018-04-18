package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// `` accent represents Unicode data that should remain constant
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"int"`
}

type User1 struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	ID    int    `xml:"id"`
}

func userRouter(w http.ResponseWriter, r *http.Request) {

	ourUser := User{}
	ourUser.Name = "Shiv Anand"
	ourUser.Email = "shivakp2111@gmail.com"
	ourUser.ID = 100

	output, err := json.Marshal(&ourUser)

	if err != nil {
		fmt.Println("something went wrong!!")
	}

	fmt.Fprintln(w, string(output))

}

func userRouter2(w http.ResponseWriter, r *http.Request) {

	ourUser := User1{}
	ourUser.Name = "Shiv Anand"
	ourUser.Email = "shivakp2111@gmail.com"
	ourUser.ID = 100

	output, err := xml.Marshal(&ourUser)

	if err != nil {
		fmt.Println("somthing went wrong!!")
	}
	fmt.Fprintln(w, string(output))

}

func main() {

	fmt.Println("Starting json server")

	http.HandleFunc("/user_json", userRouter)
	http.HandleFunc("/user_xml", userRouter2)

	// Listen on port 8080
	http.ListenAndServe(":8080", nil)

}
