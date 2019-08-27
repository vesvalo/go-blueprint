package migrate

import (
	"database/sql"
	"fmt"
	global "github.com/Nerufa/go-blueprint/cmd"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"os"
)

var (
	argDsn, argTable string
	argLimit         int
	db               *sql.DB
	ms               *migrate.FileMigrationSource
	initCmdFn        = func(cmd *cobra.Command, _ []string) (re error) {
		var e error
		defer func() {
			recover()
			re = e
		}()
		ms = &migrate.FileMigrationSource{
			Dir: global.Slave.WorkDir() + "/migrations",
		}
		migrate.SetTable(argTable)
		db, e = sql.Open("postgres", argDsn)
		if e != nil {
			return e
		}
		return nil
	}
	cmdUp = &cobra.Command{
		Use:           "up",
		Short:         "Migrates the database to the most recent version available",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			if e := initCmdFn(nil, nil); e != nil {
				fmt.Printf("Failed: %v\n", e.Error())
				os.Exit(1)
			}
			n, err := migrate.ExecMax(db, "postgres", ms, migrate.Up, argLimit)
			if err != nil {
				fmt.Printf("Failed: %v\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("Applied %d migrations!\n", n)
		},
	}
	cmdDown = &cobra.Command{
		Use:           "down",
		Short:         "Undo a database migration",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			if e := initCmdFn(nil, nil); e != nil {
				fmt.Printf("Failed: %v\n", e.Error())
				os.Exit(1)
			}
			n, err := migrate.ExecMax(db, "postgres", ms, migrate.Down, argLimit)
			if err != nil {
				fmt.Printf("Failed: %v\n", err.Error())
				os.Exit(1)
			}
			fmt.Printf("Applied %d migrations!\n", n)
		},
	}
	Cmd = &cobra.Command{
		Use:           "migrate",
		Short:         "SQL migration tool",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func init() {
	Cmd.PersistentFlags().StringVar(&argTable, "table", "migrations", "Table for migration history")
	Cmd.PersistentFlags().IntVar(&argLimit, "limit", 0, "Limit the number of migrations (0 = unlimited)")
	Cmd.PersistentFlags().StringVar(&argDsn, "dsn", "postgres://blueprint:insecure@localhost:5567?sslmode=disable", "DSN connection string")
	Cmd.AddCommand(cmdUp, cmdDown)
}
