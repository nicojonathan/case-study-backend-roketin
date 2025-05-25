package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "workbench_user:workbenchpwd@tcp(localhost:3306)/db_short_movie_festival")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func ConnectMongo() (client *mongo.Client, err error) {
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	return client, nil
}
