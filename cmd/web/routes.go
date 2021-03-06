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
<<<<<<< HEAD
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room/", handlers.Repo.BookRoom)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	
=======
	mux.Get("/make-reservation", handlers.Repo.Reservation)
>>>>>>> f9d80378920daa59ece0b025d455d909b50d8551
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/search", handlers.Repo.Search)
	mux.Post("/search", handlers.Repo.PostSearch)
	mux.Post("/search-json", handlers.Repo.PostSearchJson)
<<<<<<< HEAD

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)

	mux.Get("/user/logout", handlers.Repo.ShowLogout)
=======
>>>>>>> f9d80378920daa59ece0b025d455d909b50d8551

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(r chi.Router) {
		//r.Use(Auth)
		r.Get("/dashboard", handlers.Repo.AdminDashBoard)
		
		r.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		r.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		r.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		r.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)
		r.Get("/reservation/{src}/{id}/show", handlers.Repo.AdminShowReservation)
		r.Get("/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservation)
		r.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservation)
		r.Post("/reservation/{src}/{id}", handlers.Repo.AdminPostShowReservation)
		
	})

	return mux
}
