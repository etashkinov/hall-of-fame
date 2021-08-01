package persistence

func AddPersonPosition(personId int64, positionId int64, teamId int64, since Date, till Date, description string) (position PersonPosition, err error) {
	query := `
		INSERT INTO person_positions (person_id, position_id, team_id, since, till, description) 
		VALUES ($1, $2, $3, $4)`

	id, err := db.insertForId(query, personId, positionId, teamId, since, till, description)

	if err == nil {
		position, err = getPersonPosition(id)
	}

	return
}

func UpdatePersonPosition(personPositionId int64, since Date, till Date, description string) (position PersonPosition, err error) {
	query := `
	UPDATE person_positions SET since = $1, till = $2, description = $3
	WHERE id = $4; `

	db.Write.MustExec(query, personPositionId, since, till, description)

	return getPersonPosition(personPositionId)
}

func DeletePersonPosition(personPositionId int64) (err error) {
	return deletePersonDict("person_positions", personPositionId)
}

func DeletePersonPositions(personId int64) (err error) {
	return deletePersonDicts("person_positions", personId)
}

var selectPersonPositionQuery = buildSelectPersonDictQuery("person_positions", "pd.since, pd.description, pd.team_id,", "positions", "position_id")

func GetPersonPositions(personId int64) (positions []PersonPosition, err error) {
	err = selectPersonDicts(selectPersonPositionQuery, personId, &positions)
	return
}

func getPersonPosition(personPositionId int64) (position PersonPosition, err error) {
	err = selectPersonDict(selectPersonPositionQuery, personPositionId, &position)
	return
}
