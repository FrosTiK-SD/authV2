package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"frostik.com/auth/constants"
	"frostik.com/auth/handler"
	"frostik.com/auth/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	// Connect to Redis
	userRedisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv(constants.REDIS_HOST) + ":6379",
	})
	cacheClient := cache.New(&cache.Options{
		Redis:      userRedisClient,
		LocalCache: cache.NewTinyLFU(constants.REDIS_CACHING_LIMIT, constants.REDIS_CACHING_DURATION),
	})

	// Connect to MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv(constants.CONNECTION_STRING)).SetServerAPIOptions(serverAPI))
	if err != nil {
		fmt.Println(os.Getenv(constants.CONNECTION_STRING))
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
		UserRedis:       userRedisClient,
		UserCacheClient: cacheClient,
	}

	r.GET("/api/token/student/verify", handler.HandlerVerifyIdToken)
	r.GET("/api/token/invalidate_cache", handler.InvalidateCache)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
