// main.go

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	drivers = []Driver{
		{Name: "Jimmy Johnson", License: "ABC123"},
		{Name: "Howard Hills", License: "XYZ789"},
		{Name: "Craig Colbin", License: "DEF333"},
	}
	cars = []Car{
		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
)

func main() {
	// Load env file.
	godotenv.Load()

	// Database connection.
	// PostgreSQL.
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Berlin",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	// SQLite.
	// db, err := gorm.Open(sqlite.Open("car_dealer.db"), &gorm.Config{})

	capture(err)

	db.AutoMigrate(&Driver{})
	db.AutoMigrate(&Car{})

	// TODO: Control if the source are already exists.
	for index := range drivers {
		db.Create(&drivers[index])
	}

	for index := range cars {
		db.Create(&cars[index])
	}

	// Web server config and routes
	e := echo.New()
	e.GET("/healthcheck", GetHealthCheck)
	e.GET("/cars", GetCars(db))
	e.GET("/cars/:id", GetCar(db))
	// e.DELETE("/cars/{id}", DeleteCar(db))

	e.Logger.Fatal(e.Start(":8080"))
}

// TODO: Move this endpoings to a `views` folder.
// e.GET("/healthcheck", GetHealthCheck)
func GetHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

// e.GET("/cars", GetCars)
func GetCars(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var cars []Car
		db.Find(&cars)

		return c.JSON(http.StatusOK, &cars)
	}
}

// e.GET("/cars/:id", GetCar)
func GetCar(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		car_id := c.Param("id")

		var car Car
		err := db.First(&car, car_id).Error

		fmt.Println(car)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "Not Found")
		}

		return c.JSON(http.StatusOK, &car)
	}
}

// e.GET("/cars/:id", GetCar)
func GetDriver(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "To be implemented")
	}
}

func capture(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
