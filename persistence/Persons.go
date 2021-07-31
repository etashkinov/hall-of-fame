package persistence

import (
	"database/sql"
	"time"
)

func CreatePerson(name string, description string) (person Person, err error) {
	now := time.Now()
	query := `INSERT INTO persons (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *; `
	row := db.Write.QueryRow(query, name, description, now, now)
	return scanPerson(row)
}

func UpdatePerson(id int64, name string, description string) (person Person, err error) {
	query := `UPDATE persons 
						SET 
							updated_at = $1, 
							name = $2, 
							description = $3 
						WHERE id = $4 
						RETURNING *;`

	row := db.Write.QueryRow(query, time.Now(), name, description, id)

	person, err = scanPerson(row)

	return
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
