package persistence

func AddPersonExpertise(personId int64, expertiseId int64, since Date, level Level) (expertise PersonExpertise, err error) {
	query := `
		INSERT INTO person_expertises (person_id, expertise_id, since, level) 
		VALUES ($1, $2, $3, $4)`

	id, err := db.insertForId(query, personId, expertiseId, since, level)

	if err == nil {
		expertise, err = getPersonExpertise(id)
	}

	return
}

func UpdatePersonExpertise(personExpertiseId int64, since Date, level Level) (expertise PersonExpertise, err error) {
	query := `
	UPDATE person_expertises SET since = $1, level = $2 
	WHERE id = $3; `

	db.Write.MustExec(query, personExpertiseId, since, level)

	return getPersonExpertise(personExpertiseId)
}

func DeletePersonExpertise(personExpertiseId int64) (err error) {
	return deletePersonDict("person_expertises", personExpertiseId)
}

func DeletePersonExpertises(personId int64) (err error) {
	return deletePersonDicts("person_expertises", personId)
}

var selectPersonExpertiseQuery = buildSelectPersonDictQuery("person_expertises", "pd.since, pd.level,", "expertises", "expertise_id")

func GetPersonExpertises(personId int64) (expertises []PersonExpertise, err error) {
	err = selectPersonDicts(selectPersonExpertiseQuery, personId, &expertises)
	return
}

func getPersonExpertise(personExpertiseId int64) (expertise PersonExpertise, err error) {
	err = selectPersonDict(selectPersonExpertiseQuery, personExpertiseId, &expertise)
	return
}
