package main

import (
	"net/http"

	_ "github.com/AzureHoshi/learning-go-itgenius/cmd/api/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {

	// Create a router mux
	mux := chi.NewRouter()

	// Add middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// Register Swagger UI route
	mux.Get("/swagger/*", httpSwagger.WrapHandler)

	// Register the API routes under "/aoi/v1"
	mux.Route("/api/v1", func(r chi.Router) {

		// register routes
		r.Get("/", app.Home)
		r.Get("/about", app.About)
		r.Get("/demomoives", app.AllDemoMovies)

		// Authenticated routes
		r.Post("/authenticate", app.authenticate)
		r.Get("/refresh", app.refreshToken)
		r.Get("/logout", app.logout)

		r.Get("/movies", app.AllMovies)
		r.Get("/movies/{id}", app.GetMovie)
		r.Get("/genres", app.AllGenres)

		// Protected routes
		// r.With(app.authRequired).Get("/admin/movies", app.MovieCatalog)
		// r.With(app.authRequired).Get("/admin/movies/{id}", app.MovieForEdit)
		// r.With(app.authRequired).Post("/admin/movies", app.InsertMovie)
		// r.With(app.authRequired).Put("/admin/movies/{id}", app.UpdateMovie)
		// r.With(app.authRequired).Delete("/admin/movies/{id}", app.DeleteMovie)

		r.Route("/admin", func(r chi.Router) {

			// protected routes
			r.Use(app.authRequired)

			r.Get("/movies", app.MovieCatalog)
			r.Get("/movies/{id}", app.MovieForEdit)
			r.Post("/movies", app.InsertMovie)
			r.Put("/movies/{id}", app.UpdateMovie)
			r.Delete("/movies/{id}", app.DeleteMovie)
		})
	})

	return mux
}
