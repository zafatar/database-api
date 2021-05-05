// model.go

package main

import (
	"errors"

	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []Car
}

type Car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

// TODO: Refactor model classes as repository-like structure
type Product struct {
	gorm.Model
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) getProduct(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) updateProduct(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) deleteProduct(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) createProduct(db *gorm.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *gorm.DB, start, count int) ([]Product, error) {
	return nil, errors.New("Not implemented")
}
