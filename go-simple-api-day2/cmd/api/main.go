package main

import (
	"backend/cmd/internal/repository"
	"backend/cmd/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

// สร้างตัวแปรกำหนด port ที่จะใช้
const port = 8080

// สร้าง application struct สำหรับเก็บ config และ database connection
type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

func main() {

	// Set application config
	var app application
	app.Domain = "example.com"

	flag.StringVar(&app.DSN, "dsn", "host=192.168.1.108 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", "Postgres Data Source Name")

	// Parse the command-line arguments for JWT
	flag.StringVar(&app.JWTSecret, "jwt-secret", "supersecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.APIKey, "api-key", "mysupersecret", "API key")

	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}

	// close the database connection
	defer app.DB.Connection().Close()

	// สร้างตัวแปร auth และกำหนดค่าให้กับ issuer, audience, secret, token expiry, refresh expiry, cookie domain, cookie path, และ cookie name
	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  app.CookieDomain,
		CookiePath:    "/",
		CookieName:    "__Host-refresh-token",
	}

	// Start the server
	fmt.Println("Starting server on ", app.Domain)
	log.Printf("Starting server on port %d", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
