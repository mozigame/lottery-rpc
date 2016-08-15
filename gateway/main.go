package main

import (
	"lottery/gateway/controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {

	router := mux.NewRouter()

	account := controllers.Account{}

	router.HandleFunc("/account", account.Index).Methods("GET")
	router.HandleFunc("/account", account.Create).Methods("POST")

	n := negroni.New(negroni.NewLogger())
	//	n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router)

	n.Run(":3000")
}
