package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"frostik.com/auth/constants"
	"frostik.com/auth/handler"
	"frostik.com/auth/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/api/option"
)

var config = &firebase.Config{ProjectID: os.Getenv("OATH_PROJECT_ID")}
var app, err = firebase.NewApp(context.Background(), config, option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT"))))

func main() {
	r := gin.Default()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Connect to MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv(constants.CONNECTION_STRING)).SetServerAPIOptions(serverAPI))
	if err != nil {
		log.Fatalf("Unable to Connect to MongoDB: %v\n", err)
	} else {
		log.Println("Connected to MongoDB")
	}

	r.Use(cors.New(util.DefaultCors()))

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	handler := &handler.Handler{
		Auth:        client,
		MongoClient: mongoClient,
	}

	r.GET("/api/token/verify", handler.HandlerVerifyIdToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}

	r.Run(port) // listen and serve on 0.0.0.0:8080
}
