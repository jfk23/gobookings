package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jfk23/gobookings/internal/config"
	"github.com/jfk23/gobookings/internal/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSulf)
	mux.Use(LoadSession)


	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-room", handlers.Repo.Generals)
	mux.Get("/majors-room", handlers.Repo.Majors)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/search", handlers.Repo.Search)
	mux.Post("/search", handlers.Repo.PostSearch)
	mux.Post("/search-json", handlers.Repo.PostSearchJson)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
