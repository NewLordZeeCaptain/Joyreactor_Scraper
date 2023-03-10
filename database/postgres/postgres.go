package postgres

import (
	"Projects/models"
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InitDB() error {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "db_name",
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return err
	}

	models := []interface{}{
		(*models.User)(nil),
		(*models.Post)(nil),
	}

	for model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{Temp: true})
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertUser(user *models.User) {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
	defer db.Close()

	_, err := db.Model(user).Insert()
	if err != nil {
		panic(err)
	}
}
