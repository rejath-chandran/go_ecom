package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	DB   string
	PORT string
	log  *log.Logger
	Con  *mongo.Database
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	logger := log.New(os.Stdout, "", log.Ltime|log.Ldate)
	app := App{
		log: logger,
		PORT: os.Getenv("PORT"),
		DB: os.Getenv("DB_URL"),
	}


	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(app.DB))
	if err != nil {
		panic(err)
	}
	app.Con=client.Database("ecom")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	srv := app.Routes()
	srv.Listen(app.PORT)
	
	
}
