package persistence

func AddPersonSkill(personId int64, skillId int64, since Date, level Level) (skill PersonSkill, err error) {
	query := `
		INSERT INTO person_skills (person_id, skill_id, since, level) 
		VALUES ($1, $2, $3, $4)`

	id, err := db.insertForId(query, personId, skillId, since, level)

	if err == nil {
		skill, err = getPersonSkill(id)
	}

	return
}

func UpdatePersonSkill(personSkillId int64, since Date, level Level) (skill PersonSkill, err error) {
	query := `
	UPDATE person_skills SET since = $1, level = $2 
	WHERE id = $3; `

	db.Write.MustExec(query, personSkillId, since, level)

	return getPersonSkill(personSkillId)
}

func DeletePersonSkill(personSkillId int64) (err error) {
	return deletePersonDict("person_skills", personSkillId)
}

func DeletePersonSkills(personId int64) (err error) {
	return deletePersonDicts("person_skills", personId)
}

var selectPersonSkillQuery = buildSelectPersonDictQuery("person_skills", "pd.since, pd.level,", "skills", "skill_id")

func GetPersonSkills(personId int64) (skills []PersonSkill, err error) {
	err = selectPersonDicts(selectPersonSkillQuery, personId, &skills)
	return
}

func getPersonSkill(personSkillId int64) (skill PersonSkill, err error) {
	err = selectPersonDict(selectPersonSkillQuery, personSkillId, &skill)
	return
}
