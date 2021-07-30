package persistence

import (
	"database/sql"
	"fmt"
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

	query := `INSERT INTO persons (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *; `
	row := db.Write.QueryRow(query, person.Name, person.Description, person.CreatedAt, person.UpdatedAt)
	return scanPerson(row)
}

func UpdatePerson(id int64, name string, description string) (person Person, err error) {
	index := 1
	query := "UPDATE persons SET updated_at = $1"

	args := []interface{}{time.Now()}

	if len(name) > 0 {
		query, index = addSetter(query, "name", index)
		args = append(args, name)
	}

	if len(description) > 0 {
		query, index = addSetter(query, "description", index)
		args = append(args, description)
	}

	args = append(args, id)

	query = fmt.Sprintf("%s WHERE id = $%d RETURNING *; ", query, index+1)
	row := db.Write.QueryRow(query, args...)

	person, err = scanPerson(row)

	return
}

func addSetter(input string, name string, currentIndex int) (query string, index int) {
	index = currentIndex + 1
	return fmt.Sprintf("%s, %s = $%d", input, name, index), index
}

func GetPerson(id int64) (person Person, err error) {
	query := `SELECT * FROM persons WHERE id = $1;`

	row := db.Read.QueryRow(query, id)
	person, err = scanPerson(row)

	return
}

func GetPersons() (persons []Person, err error) {
	query := `SELECT * FROM persons;`

	rows, err := db.Read.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		person, _ := scanPersons(rows)
		persons = append(persons, person)
	}
	err = rows.Err()
	return
}

func scanPersons(row *sql.Rows) (person Person, err error) {
	err = row.Scan(&person.Id, &person.Name, &person.Description, &person.CreatedAt, &person.UpdatedAt)
	return
}

func scanPerson(row *sql.Row) (person Person, err error) {
	err = row.Scan(&person.Id, &person.Name, &person.Description, &person.CreatedAt, &person.UpdatedAt)
	return
}
