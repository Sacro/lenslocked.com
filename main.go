package main

import (
	"fmt"
	"net/http"

	"github.com/Sacro/lenslocked.com/controllers"
	"github.com/Sacro/lenslocked.com/models"
	"github.com/gorilla/mux"
)

const (
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("dbname=%s sslmode=disable", dbname)
	us, err := models.NewUserService(psqlInfo)
	must(err)
	defer us.Close()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.HomeView).Methods("GET")
	r.Handle("/contact", staticC.ContactView).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
