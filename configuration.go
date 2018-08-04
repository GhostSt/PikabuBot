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
}

// Parses configuration file and sets it to Registry
func (r *Registry) loadConfig() error {
	file, err := yaml.ReadFile("resources/config.yml")

	if (err != nil) {
		return err
	}

	r.config = file

	return nil
}

func CreateRegistry() *Registry {
	return &Registry{}
}

// Sets up application and initialize Registry
func (r *Registry) setup() error {
	err := r.loadConfig()

	if err != nil {
		return err
	}

	err = r.setupDatabase()

	if err != nil {
		return err
	}

	r.env = &env{}

	return nil
}

// Sets up connection to database and sets it to Registry and import initial database schema
func (r *Registry) setupDatabase() (error) {
	database, err := r.config.Get("database.name")

	if err != nil {
		panic(err)
	}
	username, err := r.config.Get("database.username")

	if err != nil {
		panic(err)
	}
	password, err := r.config.Get("database.password")

	if err != nil {
		panic(err)
	}

	mysql := fmt.Sprintf("%s:%s@/%s", username, password, database)
	db, _ := sql.Open("mysql", mysql)

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	r.db = db

	return nil
}