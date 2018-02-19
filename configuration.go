package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/kylelemons/go-gypsy/yaml"
	"errors"
	"os"
	"io/ioutil"
)

type registry struct {
	config *yaml.File
	db     *sql.DB
}

// Parses configuration file and sets it to Registry
func loadConfig(registry *registry) {
	file, err := yaml.ReadFile("resources/config.yml")

	if (err != nil) {
		panic(err)
	}

	registry.config = file
}

// Sets up application and initialize Registry
func setup() *registry {
	registry := &registry{}

	loadConfig(registry)

	setupDatabase(registry)

	return registry
}

// Sets up connection to database and sets it to Registry and import initial database schema
func setupDatabase(registry *registry) (error) {
	version, err := registry.config.Get("database.version")

	if err != nil {
		return errors.New("Database version doesn't set in configuration")
	}

	path, err := registry.config.Get("database.path")

	if err != nil {
		return errors.New("Database path doesn't set in configuration")
	}

	db, err := sql.Open(version, path)

	if err != nil {
		return errors.New("Database schema doesn't set in configuration")
	}

	migration_file, err := registry.config.Get("database.schema")

	if err != nil {
		panic(err)
	}

	if _, err = os.Stat(path); os.IsNotExist(err) {
		data, err := ioutil.ReadFile(migration_file)

		if err != nil {
			panic(err)
		}

		_, err = db.Exec(string(data))

		if err != nil {
			panic(err)
		}
	}

	registry.db = db

	return nil
}