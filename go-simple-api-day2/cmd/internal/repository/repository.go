package repository

import (
	"backend/cmd/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	// get User by email
	GetUserByEmail(email string) (*models.User, error)
	// get User by ID
	GetUserByID(id int) (*models.User, error)

	// get all movies
	AllMovies() ([]*models.Movie, error)
	// get one movie by ID
	OneMovie(id int) (*models.Movie, error)
	// Edit one movie by ID
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
}
