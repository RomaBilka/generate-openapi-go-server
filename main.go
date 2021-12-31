/*
 * Test REST API
 *
 * Task №3
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	ItemApiService := openapi.NewItemApiService()
	ItemApiController := openapi.NewItemApiController(ItemApiService)

	ItemsApiService := openapi.NewItemsApiService()
	ItemsApiController := openapi.NewItemsApiController(ItemsApiService)

	router := openapi.NewRouter(ItemApiController, ItemsApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}