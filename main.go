package main

import (
	"fmt"
	"os"

	"frostik.com/auth/constants"
	"frostik.com/auth/controller"
	"frostik.com/auth/handler"
	"frostik.com/auth/util"
	"github.com/FrosTiK-SD/mongik"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mongikClient := mongik.NewClient(os.Getenv(constants.CONNECTION_STRING), constants.CACHING_DURATION)

	// Initialie default JWKs
	defaultJwkSet, jwkSetRetrieveError := controller.GetJWKs(mongikClient.CacheClient, true)
	if jwkSetRetrieveError != nil {
		fmt.Println("Error retrieving JWKs")
	}

	r.Use(cors.New(util.DefaultCors()))

	handler := &handler.Handler{
		MongikClient: mongikClient,
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
