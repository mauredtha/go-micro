package api

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// App struct for new api
type App struct {
	log *log.Logger
	mux *httprouter.Router
}

// Handler type as standard http.Handle
type Handler func(http.ResponseWriter, *http.Request)

// Ctx type for encapsulated context key
type Ctx string

// Handle associates a httprouter Handle function with an HTTP Method and URL pattern.
func (a *App) Handle(method, url string, h Handler) {

	fn := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), Ctx("ps"), ps)

		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Token")
		header.Add("Content-Type", "application/json; charset=utf-8")

		h(w, r.WithContext(ctx))
	}

	a.mux.Handle(method, url, fn)
}

// HandleCors and OPTIONS response
func (a *App) HandleCors() {
	a.mux.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Add("Access-Control-Allow-Origin", "*")
			header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT")
			header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Token")
			header.Add("Content-Type", "application/json; charset=utf-8")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
}

// ServeHTTP implements the http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// NewApp for create new api
func NewApp(log *log.Logger) *App {
	return &App{log: log, mux: httprouter.New()}
}
