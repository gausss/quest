package main

import (
	"database/sql"
	"log"
	"net/http"
)

type App struct {
	Router *http.ServeMux
	DB     *sql.DB
}

func (app *App) Init(name string, password string) {
	app.Router = http.NewServeMux()
	app.initializeRoutes()
}

func (app *App) Run(port string) {
	log.Printf("Serving on %s", port)
	http.ListenAndServe(port, app.Router)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println(id)
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("GET /quest/{id}", app.getProduct)
}
