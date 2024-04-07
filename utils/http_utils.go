package utils

import (
	"fmt"
	"log"
	"net/http"
)

/*
GetRequestQuery is a helper function that abstracts the process of getting a request query form the url
*/
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

/*
Send sends the given bytes to the client with the corresponding content type
*/
func Send(writer http.ResponseWriter, data *[]byte, contentType string) {
	writer.Header().Set("Content-Type", contentType)
	_, err := writer.Write(*data)
	if err != nil {
		log.Println("Failed to send data")
	}
}
