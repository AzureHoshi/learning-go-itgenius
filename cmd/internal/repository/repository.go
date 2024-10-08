package repository

import (
	"database/sql"

	"github.com/AzureHoshi/learning-go-itgenius/cmd/internal/models"
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

	// get all genres
	AllGenres() ([]*models.Genre, error)
	// insert a new movie
	InsertMovie(m models.Movie) (int, error)
	// update a movie
	UpdateMovie(m models.Movie) error
	//  update a movie with genres
	UpdateMovieGenres(id int, genresIDs []int) error
	// delete a movie
	DeleteMovie(id int) error
}
