package bot

import (
	cli2 "github.com/urfave/cli"
	"fmt"
	"os"
	"log"
)

var registry *Registry

func init() {
	registry = CreateRegistry()
	registry.setup()
}

func main() {
	app := cli2.NewApp()
	app.EnableBashCompletion = true
	app.Usage = "PikabuBot - tool to collect posts"
	app.Version = "0.0.1"
	app.Commands = []cli2.Command{
		{
			Name:  "migrations",
			Usage: "Migration tool",
			Subcommands: []cli2.Command{
				{
					Name:  "migrate",
					Usage: "Rolls migrations",
					Flags: []cli2.Flag{
						cli2.IntFlag{
							Name: "count",
							Usage: "Rolls certain quantity of migration",
						},
					},
					Action: func(c *cli2.Context) error {
 						fmt.Println(c.Args())
						migrationManager := CreateMigrationManager()
						migrationManager.Migrate()

						return nil
					},
				},
				{
					Name:  "rollback",
					Usage: "Rollbacks migrations",
					Flags: []cli2.Flag{
						cli2.IntFlag{
							Name: "count, c",
							Usage: "Rollbacks certain quantity of migration",
						},
					},
					Action: func(c *cli2.Context) error {
						migrationManager := CreateMigrationManager()
						migrationManager.Rollback()

						return nil
					},
				},
			},
			BashComplete: func(context *cli2.Context) {
				subCommands := []string{"migrate", "rollback"}

				autoCompleteCommands(subCommands)
			},
		},
		{
			Name:  "parser",
			Usage: "Manage parser to parse content and manage data",
			Subcommands: []cli2.Command{
				{
					Name:  "start",
					Usage: "Starts parser",
					Action: func(c *cli2.Context) error {
						println("Starts")

						return nil
					},
				},
				{
					Name:  "stop",
					Usage: "Stops parser",
					Action: func(c *cli2.Context) error {
						println("Stops parser")

						return nil
					},
				},
			},
			BashComplete: func(context *cli2.Context) {
				subCommands := []string{"start", "stop"}

				autoCompleteCommands(subCommands)
			},
		},
	}
	app.BashComplete = func(context *cli2.Context) {
		subCommands := []string{"parser", "migrations"}

		autoCompleteCommands(subCommands)
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func autoCompleteCommands(commands []string) {
	for _, command := range commands {
		fmt.Println(command)
	}
}
