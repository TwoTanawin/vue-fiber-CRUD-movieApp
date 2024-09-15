package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Movie struct
type Movie struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

var db *gorm.DB

func main() {
	// Get database details from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Construct the database connection string
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Connect to database
	var err error
	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Movie{})

	app := fiber.New()

	// Use CORS for frontend communication
	app.Use(cors.New())

	// Routes
	app.Get("/movies", GetMovies)
	app.Get("/movies/:id", GetMovie)
	app.Post("/movies", CreateMovie)
	app.Put("/movies/:id", UpdateMovie)
	app.Delete("/movies/:id", DeleteMovie)

	log.Fatal(app.Listen(":3000"))
}

// Get all movies
func GetMovies(c *fiber.Ctx) error {
	var movies []Movie
	db.Find(&movies)
	return c.JSON(movies)
}

// Get single movie
func GetMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie Movie
	db.First(&movie, id)
	if movie.ID == 0 {
		return c.Status(404).SendString("Movie not found")
	}
	return c.JSON(movie)
}

// Create a new movie
func CreateMovie(c *fiber.Ctx) error {
	movie := new(Movie)
	if err := c.BodyParser(movie); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&movie)
	return c.JSON(movie)
}

// Update movie
func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie Movie
	db.First(&movie, id)
	if movie.ID == 0 {
		return c.Status(404).SendString("Movie not found")
	}
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	db.Save(&movie)
	return c.JSON(movie)
}

// Delete movie
func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie Movie
	db.First(&movie, id)
	if movie.ID == 0 {
		return c.Status(404).SendString("Movie not found")
	}
	db.Delete(&movie)
	return c.SendString("Movie deleted")
}
