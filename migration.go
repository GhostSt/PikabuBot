package main

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "database/sql"
	_ "github.com/golang-migrate/migrate/source/file"
	"fmt"
)
type MigrationManager struct {
}

func CreateMigrationManager() *MigrationManager {
	return &MigrationManager{}
}

func (migrationManager *MigrationManager) Migrate() error {
	m, err := migrationManager.createMigration()

	if err != nil {
		panic(err)
	}

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		panic(err)
	}

	fmt.Println("Migrations has been sucessfull rolled")

	return nil
}

func (migrationManager *MigrationManager) Rollback() error {
	m, err := migrationManager.createMigration()

	if err != nil {
		panic(err)
	}

	err = m.Down()

	if err != nil && err.Error() != "no change" {
		panic(err)
	}

	fmt.Println("Migrations has been sucessfull rollbacked")

	return nil
}

func (migrationManager *MigrationManager) createMigration() (*migrate.Migrate, error) {
	driver, _ := mysql.WithInstance(registry.db, &mysql.Config{})

	pathToMigrations, err := registry.config.Get("migrations.path")

	if err != nil {
		return nil, err
	}

	// TODO: move path to file to configuration
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", pathToMigrations),
		"mysql",
		driver)

	if err != nil {
		return nil, err
	}

	return m, nil
}

