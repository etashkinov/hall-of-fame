package persistence

func AddPersonAchievement(personId int64, achievementId int64, since Date, description string) (achievement PersonAchievement, err error) {
	query := `
		INSERT INTO person_achievements (person_id, achievement_id, since, description) 
		VALUES ($1, $2, $3, $4)`

	id, err := db.insertForId(query, personId, achievementId, since, description)

	if err == nil {
		achievement, err = getPersonAchievement(id)
	}

	return
}

func UpdatePersonAchievement(personAchievementId int64, since Date, description string) (achievement PersonAchievement, err error) {
	query := `
	UPDATE person_achievements SET since = $1, description = $2 
	WHERE id = $3; `

	db.Write.MustExec(query, personAchievementId, since, description)

	return getPersonAchievement(personAchievementId)
}

func DeletePersonAchievement(personAchievementId int64) (err error) {
	return deletePersonDict("person_achievements", personAchievementId)
}

func DeletePersonAchievements(personId int64) (err error) {
	return deletePersonDicts("person_achievements", personId)
}

var selectPersonAchievementQuery = buildSelectPersonDictQuery("person_achievements", "pd.since, pd.description,", "achievements", "achievement_id")

func GetPersonAchievements(personId int64) (achievements []PersonAchievement, err error) {
	err = selectPersonDicts(selectPersonAchievementQuery, personId, &achievements)
	return
}

func getPersonAchievement(personAchievementId int64) (achievement PersonAchievement, err error) {
	err = selectPersonDict(selectPersonAchievementQuery, personAchievementId, &achievement)
	return
}
