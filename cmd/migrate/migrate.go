package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	// "entgo.io/ent/dialect/sql/schema"

	migrate "github.com/fbngrm/around-home/ent/migrate"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	// _ "ariga.io/atlas/sql/postgres"
	// _ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"

	"github.com/fbngrm/around-home/pkg/database"
)

var (
	flagDSN = flag.Bool("dsn", false, "print the DSN connection string")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	dsn := database.BuildDSN("postgres", database.Config{})

	fmt.Println(dsn)
	if *flagDSN {
		return
	}

	dir, err := sqltool.NewGolangMigrateDir("ent/migrations")
	if err != nil {
		log.Fatalf("could not create migration directory: %v", err)
	}
	// Write migration diff.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeInspect),
		schema.WithDialect(dialect.Postgres),
	}

	err = migrate.NamedDiff(ctx, dsn, "match", opts...)
	if err != nil {
		log.Fatalf("could not generate migration files: %v", err)
	}
}
