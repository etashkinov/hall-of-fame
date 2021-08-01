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

func buildSelectPersonDictQuery(table string, params string, dictTable string, referenceColumn string) string {
	return fmt.Sprintf(`
	SELECT pd.id, pd.person_id, %s d.id as dict_id, d.name, d.description
	FROM %s pd
	JOIN %s d ON d.id = pd.%s
	`, params, table, dictTable, referenceColumn)
}

func deletePersonDict(dictType string, personDictId int64) (err error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1;`, dictType)

	_, err = db.Write.Exec(query, personDictId)

	return
}

func deletePersonDicts(dictType string, personId int64) (err error) {
	query := fmt.Sprintf(`DELETE FROM %s WHERE personId = $1;`, dictType)

	_, err = db.Write.Exec(query, personId)

	return
}

func selectPersonDicts(query string, personId int64, dest interface{}) error {
	return db.Read.Select(dest, fmt.Sprintf(`%s WHERE pd.person_id = $1`, query), personId)
}

func selectPersonDict(query string, personDictId int64, dest interface{}) error {
	return db.Read.Get(dest, fmt.Sprintf(`%s WHERE pd.id = $1`, query), personDictId)
}
