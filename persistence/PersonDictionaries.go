package persistence

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

const SELECT_PERSON_SKILL = `
SELECT ps.id, ps.person_id, ps.since, ps.level, s.id as dict_id, s.name, s.description
FROM person_skills ps
JOIN skills s ON s.id = ps.skill_id
`

func AddPersonSkill(personId int64, skillId int64, since Date, level Level) (skill PersonSkill, err error) {
	query := `
		INSERT INTO person_skills (person_id, skill_id, since, level) 
		VALUES ($1, $2, $3, $4) RETURNING id; `

	row := db.Write.QueryRow(query, personId, skillId, since, level)

	var id int64
	err = row.Scan(&id)
	if err != nil {
		return
	}

	skill, err = getPersonSkill(id)
	return
}

func GetPersonSkills(personId int64) (skills []PersonSkill, err error) {
	query := fmt.Sprintf(`%s WHERE ps.person_id = $1`, SELECT_PERSON_SKILL)
	err = db.Read.Select(&skills, query, personId)
	return
}

func getPersonSkill(personSkillId int64) (skill PersonSkill, err error) {
	query := fmt.Sprintf(`%s WHERE ps.id = $1`, SELECT_PERSON_SKILL)
	err = db.Read.Get(&skill, query, personSkillId)
	return
}

func (date Date) Value() (driver.Value, error) {
	return date.Format(), nil
}

func (date *Date) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		return nil
	}
	if str, err := driver.String.ConvertValue(value); err == nil {
		// if this is a bool type
		if v, ok := str.(string); ok {
			dayPart := strings.Fields(v)[0]
			err = date.Parse(dayPart)
			return err
		}
	}
	// otherwise, return an error
	return fmt.Errorf("Failed to scan %s date", value)
}
