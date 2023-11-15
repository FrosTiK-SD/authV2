package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"frostik.com/auth/constants"
	"frostik.com/auth/controller"
	"frostik.com/auth/handler"
	"frostik.com/auth/util"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	// Initialize BigCache
	cacheClient, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(constants.CACHING_DURATION))

	// Initialie default JWKs
	defaultJwkSet, jwkSetRetrieveError := controller.GetJWKs(cacheClient, true)
	if jwkSetRetrieveError != nil {
		fmt.Println("Error retrieving JWKs")
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

	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	handler := &handler.Handler{
		MongoClient:     mongoClient,
		UserCacheClient: cacheClient,
		JwkSet:          defaultJwkSet,
	}

	r.GET("/api/token/student/verify", handler.HandlerVerifyStudentIdToken)
	r.GET("/api/token/invalidate_cache", handler.InvalidateCache)

	port := "" + os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
