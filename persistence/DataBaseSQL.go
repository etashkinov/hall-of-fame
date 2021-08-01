package persistence

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type dbConfig struct {
	Hostname string
	Name     string
	Username string
	Password string
	Port     string
}

type infoDatabase struct {
	Read  dbConfig
	Write dbConfig
}

// Host databases to work
var (
	db database
	// dbOtherDB       Database
)

// Nodes read and write in database
type database struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

func init() {
	var infoDB infoDatabase
	viper.SetConfigFile("config.json")
	_ = viper.ReadInConfig()

	_ = mapstructure.Decode(viper.GetStringMap("Databases.PostgreSQL"), &infoDB)
	_ = db.upConnection(&infoDB)
}

// Up new mysql database connection

func connect(config *dbConfig) (connection *sqlx.DB, err error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s  password=%s sslmode=disable", config.Hostname, config.Port, config.Username, config.Password)
	connection, err = sqlx.Connect("postgres", connectString)
	connection.SetConnMaxLifetime(time.Second * 10)
	return
}

func (db *database) upConnection(info *infoDatabase) (err error) {
	db.Read, err = connect(&info.Read)
	if err != nil {
		return
	}

	db.Write, err = connect(&info.Write)

	db.migrateDB()

	return
}

func (db *database) migrateDB() {
	driver, err := postgres.WithInstance(db.Write.DB, &postgres.Config{})
	if err != nil {
		fmt.Printf("Failed to migrate: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)

	if err != nil {
		fmt.Printf("Failed to migrate: %s", err)
	}

	m.Steps(2)
}

func (db database) insertForId(query string, args ...interface{}) (int64, error) {
	row := db.Write.QueryRow(fmt.Sprintf("%s RETURNING id;", query), args...)

	return getId(row)
}

func getId(row *sql.Row) (id int64, err error) {
	err = row.Scan(&id)
	return
}

func (db database) insertForResult(result *interface{}, query string, args ...interface{}) error {
	row := db.Write.QueryRowx(fmt.Sprintf("%s RETURNING *;", query), args)
	return row.Scan(result)
}
