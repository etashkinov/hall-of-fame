package persistence

import (
	"database/sql"
	"fmt"
)

func CreateDictionary(dictType string, name string, description string) (dict Dictionary, err error) {
	query := fmt.Sprintf("INSERT INTO %s (name, description) VALUES ($1, $2) RETURNING *; ", dictType)
	row := db.Write.QueryRow(query, name, description)
	return scanDictionary(row)
}

func UpdateDictionary(dictType string, id int64, name string, description string) (person Dictionary, err error) {
	query := fmt.Sprintf("UPDATE %s SET name = $1, description = $2 WHERE id = $3 RETURNING *;", dictType)
	row := db.Write.QueryRow(query, name, description, id)
	return scanDictionary(row)
}

func GetDictionaries(dictType string) (dicts []Dictionary, err error) {
	query := fmt.Sprintf("SELECT * FROM %s;", dictType)
	rows, err := db.Read.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		dict, _ := scanDictionaries(rows)
		dicts = append(dicts, dict)
	}
	err = rows.Err()
	return
}

func DeleteDictionary(dictType string, id int64) (dict Dictionary, err error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING *;", dictType)
	row := db.Write.QueryRow(query, id)
	return scanDictionary(row)
}

func scanDictionaries(rows *sql.Rows) (dict Dictionary, err error) {
	err = rows.Scan(&dict.Id, &dict.Name, &dict.Description)
	return
}

func scanDictionary(row *sql.Row) (dict Dictionary, err error) {
	err = row.Scan(&dict.Id, &dict.Name, &dict.Description)
	return
}
