CREATE TABLE persons (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE TABLE skills (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT
);
CREATE TABLE person_skills (
  id SERIAL PRIMARY KEY,
  person_id BIGINT REFERENCES persons(id) NOT NULL,
  skill_id BIGINT REFERENCES skills(id) NOT NULL,
  since DATE NOT NULL,
  level int NOT NULL
);
CREATE TABLE achievements (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT
);
CREATE TABLE person_achievements (
  id SERIAL PRIMARY KEY,
  person_id BIGINT REFERENCES persons(id) NOT NULL,
  achievements BIGINT REFERENCES achievements(id) NOT NULL,
  since DATE NOT NULL,
  description TEXT
);
CREATE TABLE expertises (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT
);
CREATE TABLE person_expertises (
  id SERIAL PRIMARY KEY,
  person_id BIGINT REFERENCES persons(id) NOT NULL,
  expertise_id BIGINT REFERENCES expertises(id) NOT NULL,
  since DATE NOT NULL,
  level int NOT NULL
);
CREATE TABLE positions (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT
);
CREATE TABLE teams (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT
);
CREATE TABLE person_positions (
  id SERIAL PRIMARY KEY,
  person_id BIGINT REFERENCES persons(id) NOT NULL,
  position_id BIGINT REFERENCES positions(id) NOT NULL,
  team_id BIGINT REFERENCES teams(id) NOT NULL,
  since DATE NOT NULL,
  till DATE,
  description TEXT
);