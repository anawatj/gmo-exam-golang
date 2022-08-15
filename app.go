package main

import (
	"API/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname, host string) {
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}
func (a *App) initializeRoutes() {
	var route routes.Route
	route.DB = a.DB
	a.Router.HandleFunc("/api/usergroup", route.GetAllUserGroup).Methods("GET")
	a.Router.HandleFunc("/api/register", route.RegisterUser).Methods("POST")
	a.Router.HandleFunc("/api/quiz", route.GetAllQuestion).Methods("GET")
	a.Router.HandleFunc("/api/load", route.LoadSaveAnswer).Methods("POST")
	a.Router.HandleFunc("/api/save", route.SaveAnswers).Methods("POST")
	a.Router.HandleFunc("/api/submit", route.SubmitUserAnswer).Methods("POST")
}
func (a *App) Run(addr string) {

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})
	log.Fatal(http.ListenAndServe(addr, corsOpts.Handler(a.Router)))
}
