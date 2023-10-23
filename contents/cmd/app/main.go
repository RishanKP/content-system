package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "content-system/contents/cmd/app/docs"
	"content-system/contents/pkg/models/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Contents *db.ContentModel
}

// @title Content Microservice
// @description API for managing contents
// @version 1.0.0
// @host localhost:9000
// @BasePath /contents
func main() {
	serverAddr := flag.String("serverAddr", "localhost", "server network address")
	serverPort := flag.Int("serverPort", 8001, "server network port")
	mongoDatabase := flag.String("mongoDatabase", os.Getenv("MONGO_DATABASE"), "mongo database")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)

	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_CLUSTER"))

	connection := options.Client().ApplyURI(mongoURI)

	client, err := mongo.NewClient(connection)
	if err != nil {
		errorLog.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}

	infoLog.Printf("Connected to Database")

	app := &Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Contents: &db.ContentModel{
			Client: client.Database(*mongoDatabase).Collection("contents"),
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errorLog,
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
