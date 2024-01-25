package main

import (
	"context"
	"fmt"
	"os"

	"frostik.com/auth/constants"
	"frostik.com/auth/controller"
	"frostik.com/auth/handler"
	"frostik.com/auth/util"
	"github.com/FrosTiK-SD/mongik"
	mongikConstants "github.com/FrosTiK-SD/mongik/constants"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mongikClient := mongik.NewClient(os.Getenv(constants.CONNECTION_STRING), &mongikModels.Config{
		Client: mongikConstants.REDIS,
		TTL:    constants.CACHING_DURATION,
		RedisConfig: &mongikModels.RedisConfig{
			URI:      os.Getenv(constants.REDIS_URI),
			Password: os.Getenv(constants.REDIS_PASSWORD),
			Username: os.Getenv(constants.REDIS_USERNAME),
		},
	})

	cacheClient, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(constants.CACHING_DURATION))

	// Initialie default JWKs
	defaultJwkSet, jwkSetRetrieveError := controller.GetJWKs(cacheClient, true)
	if jwkSetRetrieveError != nil {
		fmt.Println("Error retrieving JWKs")
	}

	r.Use(cors.New(util.DefaultCors()))

	handler := &handler.Handler{
		MongikClient: mongikClient,
		BigCache:     cacheClient,
		JwkSet:       defaultJwkSet,
	}

	r.GET("/api/token/student/verify", handler.HandlerVerifyStudentIdToken)
	r.GET("/api/token/invalidate_cache", handler.InvalidateCache)

	port := "" + os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
