package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// GetRequestQuery retrieves the query parameter with the specified key from the request URL.
//
// Parameters:
// - writer: http.ResponseWriter for writing response.
// - request: *http.Request containing the request data.
// - key: string representing the key of the query parameter to retrieve.
// Return type: *string pointing to the value of the query parameter, or nil if the parameter doesn't exist.
func GetRequestQuery(writer http.ResponseWriter, request *http.Request, key string) *string {
	queryStr := request.URL.Query().Get(key)
	if queryStr == "" {
		errorStr := fmt.Sprintf("Parameter %s doesn't exist", key)
		log.Println(errorStr)
		http.Error(writer, errorStr, http.StatusBadRequest)
		return nil
	}

	return &queryStr
}

// Send sends data as JSON with the specified content type.
//
// Parameters:
//   - writer: http.ResponseWriter for sending the response.
//   - data: interface{} containing the data to be sent.
//   - contentType: string specifying the content type of the response.
func Send(writer http.ResponseWriter, data interface{}, contentType string) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, "Server unable to marshal data to bytes", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", contentType)

	_, err = writer.Write(dataBytes)
	if err != nil {
		log.Println("Failed to send data")
	}
}

// ReadBody reads the body from the http request and unmarshals it into the provided interface.
//
// Parameters:
// - writer: the http.ResponseWriter to write the response.
// - request: the *http.Request to read the body from.
// - tfo: the interface{} to unmarshal the body into.
// Returns an error if any.
func ReadBody(writer http.ResponseWriter, request *http.Request, tfo interface{}) error {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	err = json.Unmarshal(body, &tfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

type Middleware func(next http.Handler) http.Handler

// LoadMiddlewares loads the provided middleware functions onto the given handler.
//
// Parameters:
// - handler (http.Handler): The original HTTP handler to apply the middlewares to.
// middlewares (...Middleware): The list of middleware functions to apply.
//
// Return:
// - http.Handler: The handler with all the middlewares applied.
func LoadMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
