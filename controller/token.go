package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"frostik.com/auth/constants"
	"frostik.com/auth/util"
	"github.com/allegro/bigcache/v3"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func getJWKs(cacheClient *bigcache.BigCache, noCache bool) (*jwk.Set, *string) {
	// Check if copy is there in the cache
	var jwkString string
	var jwkBytes []byte

	if !noCache {
		jwkBytes, err := cacheClient.Get(constants.GCP_JWKS)
		if err == nil {
			fmt.Println("Successfully fetched JWKs from cache")
			jwkString = string(jwkBytes)
			jwkSet, err := jwk.ParseString(jwkString)
			if err != nil {
				return nil, &constants.ERROR_PARSING_JWK
			} else {
				return &jwkSet, nil
			}
		}
	}

	// Fetch the JWKs from GoogleAPIs
	jwks, err := http.Get("https://www.googleapis.com/service_accounts/v1/jwk/securetoken@system.gserviceaccount.com")
	if err != nil {
		return nil, &constants.ERROR_FETCH_JWK
	}
	fmt.Println("Fetched JWKs from GCP")

	// Convert to bytes and them read it as a string
	jwkBytes, err = io.ReadAll(jwks.Body)
	if err != nil {
		return nil, &constants.ERROR_CONVERT_JWT_TO_BYTES
	}

	jwkString = string(jwkBytes)
	jwkSet, err := jwk.ParseString(jwkString)
	if err != nil {
		return nil, &constants.ERROR_PARSING_JWK
	}

	// Set the JWKs in the cache
	if err = cacheClient.Set(constants.GCP_JWKS, []byte(jwkString)); err == nil {
		fmt.Println("Successfully set JWKs in cache")
	}

	return &jwkSet, nil
}

func VerifyToken(cacheClient *bigcache.BigCache, idToken string, noCache bool) (*string, *string) {
	jwkSet, jwkParsingError := getJWKs(cacheClient, noCache)
	if jwkParsingError != nil {
		return nil, jwkParsingError
	}

	// Verify the token
	rawJWT, err := jwt.Parse([]byte(idToken), jwt.WithKeySet(*jwkSet))
	if err != nil {
		return nil, &constants.ERROR_TOKEN_SIGNATURE_INVALID
	}

	// Validations
	if time.Now().Sub(rawJWT.IssuedAt()) < 0 || time.Now().Sub(rawJWT.Expiration()) > 0 || rawJWT.Subject() == "" || rawJWT.Issuer() != fmt.Sprintf("https://securetoken.google.com/%s", os.Getenv(constants.FIREBASE_PROJECT_ID)) || !util.ArrayContains(rawJWT.Audience(), os.Getenv(constants.FIREBASE_PROJECT_ID)) {
		return nil, &constants.ERROR_INVALID_TOKEN
	}

	// Get the email
	email, found := rawJWT.Get("email")
	if found == false {
		return nil, &constants.ERROR_GETTING_EMAIL
	}

	emailString := fmt.Sprintf("%v", email)

	return &emailString, nil
}
