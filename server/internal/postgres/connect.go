package postgres

import (
	ent "github.com/fbngrm/around-home/ent"
	"github.com/fbngrm/around-home/pkg/database"
	_ "github.com/lib/pq"
)

func Connect(conf database.Config) (*ent.Client, error) {
	return ent.Open("postgres", database.BuildDSN("postgres", conf))
}
