package postgres

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/advenjourney/api/pkg/config"
	"github.com/cenkalti/backoff/v4"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/jackc/pgx/v4/pgxpool"

	// initialize migrate file
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *pgxpool.Pool

func InitDB(ctx context.Context, config config.Config) {
	log.Println(config.Database.DSN)
	connectToDB := func() error {
		var err error
		DB, err = pgxpool.Connect(ctx, config.Database.DSN)

		return err
	}

	err := backoff.RetryNotify(connectToDB, backoff.NewExponentialBackOff(), func(err error, duration time.Duration) {
		log.Printf("Could not connect to database: %s, ...retrying!\n", err)
	})

	if err != nil {
		log.Fatal(err)
	}
}

// Migrate uses stdlib-sql to stay compatible with migrate package
func Migrate() {
	sqldb, err := sql.Open("postgres", DB.Config().ConnString())
	if err != nil {
		log.Panic(err)
	}

	if err = sqldb.Ping(); err != nil {
		sqldb.Close()
		log.Fatal(err)
	}

	driver, _ := postgres.WithInstance(sqldb, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/postgres/migrations",
		"postgres",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		sqldb.Close()
		log.Fatal(err)
	}

	sqldb.Close()
}
