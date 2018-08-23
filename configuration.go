package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)

type Registry struct {
	config *yaml.File
	db     *sql.DB
	env    *env
	logger Logger
}

// Parses configuration file and sets it to Registry
func (r *Registry) loadConfig() {
	file, err := yaml.ReadFile("resources/config.yml")

	if (err != nil) {
		r.logger.Panic(err.Error())
	}

	r.config = file
}

func CreateRegistry() *Registry {
	return &Registry{}
}

// Sets up application and initialize Registry
func (r *Registry) setup() {
	r.logger = &DefaultLogger{}
	r.env = &env{}
	r.loadConfig()
	r.setupLogger()
	r.setupDatabase()
}

func (r *Registry) setupLogger() {
	loggerType, err := r.config.Get("logger.type")

	if err != nil {
		r.logger.Panic(err.Error())
	}

	switch loggerType {
	case "syslog":
		newLogger, err := CreateSyslogLogger()

		if err != nil {
			r.logger.Panic(err.Error())
		}

		r.logger = newLogger
	case "default":
	default:
		r.logger.Panic(fmt.Sprintf("Logger configuration is invalid. Logger type \"%s\" not found", loggerType))
	}
}

// Sets up connection to database and sets it to Registry and import initial database schema
func (r *Registry) setupDatabase() {
	database, err := r.config.Get("database.name")

	if err != nil {
		r.logger.Panic(err.Error())
	}
	username, err := r.config.Get("database.username")

	if err != nil {
		r.logger.Panic(err.Error())
	}
	password, err := r.config.Get("database.password")

	if err != nil {
		r.logger.Panic(err.Error())
	}

	mysql := fmt.Sprintf("%s:%s@/%s", username, password, database)
	db, _ := sql.Open("mysql", mysql)

	err = db.Ping()

	if err != nil {
		r.logger.Panic(err.Error())
	}

	r.db = db
}