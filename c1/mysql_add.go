package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// database driver
var database *sql.DB

type User struct {
	ID    int    "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

// create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	output, err := json.Marshal(NewUser)

	fmt.Println(string(output))

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_firstname='" + NewUser.First + "', user_lastname='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
	q, err := database.Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(q)
}

// getting user from database
func GetUser(w http.ResponseWriter, r *http.Request) {

	// mux function Vars from getting the values
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	ReadUser := User{}
	
	err := database.QueryRow("select * from users where
		user_id=?",id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First,
		&ReadUser.Last, &ReadUser.Email)
	
	switch {
		case err == sql.ErrNoRows:
			fmt.Fprintf(w, "No such user")
		
		case err != nil:
			log.Fatal(err)
		
		default:
			// marshal the json value obtained
			output, _ := json.Marshal(ReadUser)
			fmt.Fprintf(w, string(output))
	}
}

func main() {

	// connect to the database
	db, err := sql.Open("mysql", "root:21071996@tcp(127.0.0.1:3306)/social_network")

	if err != nil {
		log.Fatal(err)
		fmt.Println("database connectivity problem")
	}

	// declared globally
	database = db
	routes := mux.NewRouter()

	// call createUser function with this route
	routes.HandleFunc("/api/user/create", CreateUser)
	// fetch the user
	routes.HandleFunc("/api/user/read/{id:[0-9]+}", GetUser)
	http.Handle("/", routes)

	// listen to this port
	http.ListenAndServe(":8080", nil)
}
