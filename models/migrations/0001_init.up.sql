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
  since TIMESTAMP WITH TIME ZONE NOT NULL,
  level VARCHAR(128) NOT NULL
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
  since TIMESTAMP WITH TIME ZONE NOT NULL,
  description TEXT
);
CREATE TABLE experises (
  id SERIAL PRIMARY KEY,
  name VARCHAR(128) NOT NULL,
  description TEXT
);
CREATE TABLE person_experises (
  id SERIAL PRIMARY KEY,
  person_id BIGINT REFERENCES persons(id) NOT NULL,
  expertise_id BIGINT REFERENCES experises(id) NOT NULL,
  since TIMESTAMP WITH TIME ZONE NOT NULL,
  level VARCHAR(128) NOT NULL
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
  since TIMESTAMP WITH TIME ZONE NOT NULL,
  till TIMESTAMP WITH TIME ZONE,
  description TEXT
);