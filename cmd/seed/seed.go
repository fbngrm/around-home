package main

import (
	"context"
	"fmt"
	"log"

	ent "github.com/fbngrm/around-home/ent"
	"github.com/fbngrm/around-home/pkg/database"
	"github.com/fbngrm/around-home/pkg/materials"
	_ "github.com/lib/pq"
)

func main() {
	if err := seed(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func seed() error {
	ctx := context.Background()

	db, err := ent.Open("postgres", database.BuildDSN("postgres", database.Config{}))
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("seeding the database...")
	if err := seedDB(ctx, db); err != nil {
		return err
	}
	fmt.Println("done")
	return nil
}

func withTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			if errRoll := tx.Rollback(); errRoll != nil {
				log.Printf("error rolling back: %+v", errRoll)
			}
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}

func seedDB(ctx context.Context, db *ent.Client) error {
	return withTx(ctx, db, func(tx *ent.Tx) error {

		wood := db.Material.Create().SetMaterial(materials.Wood).SaveX(ctx)
		tiles := db.Material.Create().SetMaterial(materials.Tiles).SaveX(ctx)
		carpet := db.Material.Create().SetMaterial(materials.Carpet).SaveX(ctx)

		bulk := []*ent.PartnerCreate{
			db.Partner.Create().SetAddress("52.519065/13.406083").SetRadiusOfOperation(10000.0).SetRating(5).AddMaterials(wood, tiles, carpet),
			db.Partner.Create().SetAddress("52.519065/13.406083").SetRadiusOfOperation(10000.0).SetRating(4).AddMaterials(wood, tiles, carpet),
			db.Partner.Create().SetAddress("52.532566/13.396261").SetRadiusOfOperation(100.0).SetRating(5).AddMaterials(wood, tiles),
			db.Partner.Create().SetAddress("52.532566/13.406083").SetRadiusOfOperation(10000.0).SetRating(5).AddMaterials(carpet),
		}
		_, err := db.Partner.CreateBulk(bulk...).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}
