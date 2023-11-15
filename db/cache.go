package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"frostik.com/auth/constants"
	"github.com/allegro/bigcache/v3"
)

func getDBClusterFromKey(key string) string {
	// We store DB cache in the format CLUSTER_NAME | OPERATION | QUERY
	return strings.Split(key, " | ")[0]
}

func DBCacheSet(cacheClient *bigcache.BigCache, key string, value []byte) error {
	// Get the list of keys
	var keyStore map[string][]string
	keyStoreBytes, _ := cacheClient.Get(constants.KEY_STORE)
	if err := json.Unmarshal(keyStoreBytes, &keyStore); err != nil {
		keyStore = make(map[string][]string)
	}

	clusterName := getDBClusterFromKey(key)

	// Add it to the cluster set
	keyStore[clusterName] = append(keyStore[clusterName], key)

	// Set the key store
	keyStoreBytes, _ = json.Marshal(keyStore)
	_ = cacheClient.Set(constants.KEY_STORE, keyStoreBytes)

	fmt.Printf("%v", keyStore)

	// return caheClient.Set(key, value)
	return cacheClient.Set(key, value)
}

func DBCacheReset(cacheClient *bigcache.BigCache, clusterName string) {
	// Get the list of keys
	var keyStore map[string][]string
	keyStoreBytes, _ := cacheClient.Get(constants.KEY_STORE)
	if err := json.Unmarshal(keyStoreBytes, &keyStore); err != nil {
		keyStore = make(map[string][]string)
	}

	// Delete all the keys in the cluster
	for _, key := range keyStore[clusterName] {
		cacheClient.Delete(key)
	}

	keyStore[clusterName] = []string{}

	// Set the key store
	keyStoreBytes, _ = json.Marshal(keyStore)
	cacheClient.Set(constants.KEY_STORE, keyStoreBytes)
}
