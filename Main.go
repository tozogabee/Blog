package main

import (
	"Blog/database"
	"Blog/handlers"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	/*menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(opts); return nil })

	menu.Option("Add a new Person", 0, true, nil)
	menu.Option("Find a Person", 1, false, nil)
	menu.Option("Update a Person's information", 2, false, nil)
	menu.Option("Delete a person by ID", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}*/
	DB := database.Connect()
	handleRequests(DB)

}

func handleRequests(DB *sql.DB) {
	h := handlers.New(DB)
	//DB.Ping()
	// create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/comments", h.GetAllComments).Methods(http.MethodGet)
	myRouter.HandleFunc("/comments/create", h.CreateComment).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
