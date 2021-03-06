/*
 * Test REST API
 *
 * Task №3
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// ItemApiRouter defines the required methods for binding the api requests to a responses for the ItemApi
// The ItemApiRouter implementation should parse necessary information from the http request,
// pass the data to a ItemApiServicer to perform the required actions, then write the service results to the http response.
type ItemApiRouter interface { 
	ItemIdDelete(http.ResponseWriter, *http.Request)
	ItemIdGet(http.ResponseWriter, *http.Request)
	ItemIdPut(http.ResponseWriter, *http.Request)
	ItemPost(http.ResponseWriter, *http.Request)
}
// ItemsApiRouter defines the required methods for binding the api requests to a responses for the ItemsApi
// The ItemsApiRouter implementation should parse necessary information from the http request,
// pass the data to a ItemsApiServicer to perform the required actions, then write the service results to the http response.
type ItemsApiRouter interface { 
	ItemsGet(http.ResponseWriter, *http.Request)
}


// ItemApiServicer defines the api actions for the ItemApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ItemApiServicer interface { 
	ItemIdDelete(context.Context, int64) (ImplResponse, error)
	ItemIdGet(context.Context, int64) (ImplResponse, error)
	ItemIdPut(context.Context, int64, CreateItemRequest) (ImplResponse, error)
	ItemPost(context.Context, CreateItemRequest) (ImplResponse, error)
}


// ItemsApiServicer defines the api actions for the ItemsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ItemsApiServicer interface { 
	ItemsGet(context.Context) (ImplResponse, error)
}
