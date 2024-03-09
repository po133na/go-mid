package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Animal struct {
	ID             string `json:"id"`
	Kind_Of_Animal string `json:"kind_of_animal"`
	Kind_Of_Breed  string `json:"kind_of_breed"`
	Name           string `json:"name"`
	Age            string `json:"age"`
	Description    string `json:"description"`
}

type AnimalModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (a AnimalModel) Insert(animal *Animal) error {
	// check for ID needed here if error
	query := `
		INSERT INTO Animals (Name, Age, Description) 
		VALUES ($1, $2, $3) 
		RETURNING id, kind_of_animal, kind_of_breed
		`
	// check if its animal of Animals in case of error
	args := []interface{}{animal.Name, animal.Age, animal.Description}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&animal.ID)
}

func (a AnimalModel) Get(id int) (*Animal, error) {
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT ID, Kind_Of_Animal, Kind_Of_Breed, Name, Age, Description
		FROM Animals
		WHERE ID = $1
		`
	var animal Animal
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// again animal or Animals?
	row := a.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&animal.ID, &animal.Kind_Of_Animal, &animal.Kind_Of_Breed, &animal.Name, &animal.Age, &animal.Description)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

func (a AnimalModel) Delete(id int) error {
	// Delete a specific menu item from the database.
	query := `
		DELETE FROM Animals
		WHERE ID = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := a.DB.ExecContext(ctx, query, id)
	return err
}
