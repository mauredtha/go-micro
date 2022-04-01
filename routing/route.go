package routing

import (
	"database/sql"
	"log"
	"microservices/controllers"
	"microservices/libraries/api"
	"net/http"
)

// API handling routing
func API(db *sql.DB, log *log.Logger) http.Handler {
	app := api.NewApp(log)
	app.HandleCors()

	// Users routing
	{
		users := controllers.Users{Db: db, Log: log}
		app.Handle(http.MethodGet, "/users", users.List)
		app.Handle(http.MethodPost, "/users", users.Create)
		app.Handle(http.MethodGet, "/users/:id", users.View)
		app.Handle(http.MethodPut, "/users/:id", users.Update)
		app.Handle(http.MethodDelete, "/users/:id", users.Delete)
	}

	return app
}
