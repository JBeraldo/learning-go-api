package database

import (
	"context"
	"jberaldo/go-api/configs"
	"log"

	"github.com/go-pg/pg/v10"
)

var DB *pg.DB = Setup()

func Setup() *pg.DB {
	log.Println(configs.Envs.DBUser)
	db := pg.Connect(&pg.Options{
		Addr:     configs.Envs.DBAddress,
		User:     configs.Envs.DBUser,
		Password: configs.Envs.DBPassword,
		Database: configs.Envs.DBName,
	})

	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Connected")

	return db
}
