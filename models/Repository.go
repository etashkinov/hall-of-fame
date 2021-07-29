package models

import (
	"time"
)

func CreatePerson(name string) (person Person, err error) {
	now := time.Now()
	person = Person{
		Name:        name,
		Description: "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	var personId int64
	query := `INSERT INTO persons (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id; `
	err = db.Write.QueryRow(query, person.Name, person.Description, person.CreatedAt, person.UpdatedAt).Scan(&personId)
	if err != nil {
		return
	}

	person.Id = personId
	return
}

func GetPerson(id int64) (person Person, err error) {
	query := `SELECT 
		id, 
		name, 
		description, 
		created_at, 
		updated_at FROM persons WHERE id = $1;`

	row := db.Read.QueryRow(query, id)
	err = row.Scan(&person.Id, &person.Name, &person.Description, &person.CreatedAt, &person.UpdatedAt)
	if err != nil {
		return
	}

	return
}
