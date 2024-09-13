package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AzureHoshi/learning-go-itgenius/cmd/internal/repository"
	"github.com/AzureHoshi/learning-go-itgenius/cmd/internal/repository/dbrepo"
	"github.com/joho/godotenv"
)

// สร้างตัวแปรกำหนด port ที่จะใช้
const port = 3000

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

// @title Movie API with GO and Postgres
// @version 1.0
// @description This is a sample movie API with GO and Postgres
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {

	// Set application config
	var app application
	app.Domain = "example.com"

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ** สร้างตัวแปร DSN สำหรับเก็บค่า connection string ของฐานข้อมูล
	// ? old code
	// flag.StringVar(&app.DSN, "dsn", "host=192.168.1.108 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable", "Postgres Data Source Name")

	// ? new code
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s timezone=%s connect_timeout=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_SSLMODE"), os.Getenv("DB_TIMEZONE"),
		os.Getenv("DB_CONNECT_TIMEOUT"))

	// อ้่างค่า DSN ไปยัง app.DSN
	app.DSN = dsn

	// ** Parse the command-line arguments for JWT
	// ? old code
	// flag.StringVar(&app.JWTSecret, "jwt-secret", "supersecret", "signing secret")
	// flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	// flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	// flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	// flag.StringVar(&app.APIKey, "api-key", "mysupersecret", "API key")

	// ? new code
	app.JWTSecret = os.Getenv("JWT_SECRET")
	app.JWTIssuer = os.Getenv("JWT_ISSUER")
	app.JWTAudience = os.Getenv("JWT_AUDIENCE")
	app.CookieDomain = os.Getenv("COOKIE_DOMAIN")
	app.Domain = os.Getenv("DOMAIN")
	app.APIKey = os.Getenv("API_KEY")

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
	err = http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), app.routes())
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
