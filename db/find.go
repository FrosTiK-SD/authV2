package db

import (
	"encoding/json"
	"fmt"

	"frostik.com/auth/constants"
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getKey(collectionName string, operation string, query bson.M) string {
	return fmt.Sprintf("%s | %s | %v", collectionName, constants.DB_FINDONE, query)
}

// Should be called after every write operation on a cluster
func resetCache(cacheClient *bigcache.BigCache, clusterName string) {
	DBCacheReset(cacheClient, clusterName)
}

func FindOne[Result any](ctx *gin.Context, mongoClient *mongo.Client, cacheClient *bigcache.BigCache, collectionName string, query bson.M, result *Result, noCache bool) {
	key := getKey(collectionName, constants.DB_FINDONE, query)
	var resultBytes []byte
	var resultInterface map[string]interface{}

	// First Check if it is present in the cache
	if !noCache {
		resultBytes, _ := cacheClient.Get(key)
		if err := json.Unmarshal(resultBytes, &result); err == nil {
			return
		}
	}

	// Query to DB
	mongoClient.Database(constants.DB).Collection(collectionName).FindOne(ctx, query).Decode(&resultInterface)

	mapstructure.Decode(resultInterface, &result)

	// Set to cache
	resultBytes, _ = json.Marshal(result)
	DBCacheSet(cacheClient, key, resultBytes)
}

func Find[Result any](ctx *gin.Context, mongoClient *mongo.Client, cacheClient *bigcache.BigCache, collectionName string, query bson.M, noCache bool) ([]Result, error) {
	key := getKey(collectionName, constants.DB_FIND, query)
	var resultBytes []byte
	var result []Result
	var resultInterface []map[string]interface{}

	// First Check if it is present in the cache
	if !noCache {
		resultBytes, _ := cacheClient.Get(key)
		if err := json.Unmarshal(resultBytes, &result); err == nil {
			return result, nil
		}
	}

	// Query to DB
	fmt.Println("Queriying the DB")
	cursor, err := mongoClient.Database(constants.DB).Collection(collectionName).Find(ctx, query)
	if err != nil {
		return nil, err
	}
	cursor.All(ctx, &resultInterface)

	mapstructure.Decode(resultInterface, &result)

	// Set to cache
	resultBytes, _ = json.Marshal(result)
	DBCacheSet(cacheClient, key, resultBytes)

	return result, nil
}
