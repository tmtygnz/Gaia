package utils

import (
	"encoding/json"
	"fmt"
	"io"
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
