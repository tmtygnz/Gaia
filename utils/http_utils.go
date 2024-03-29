package utils

import (
	"log"
	"net/http"
)

func GetRequestQuery(writer *http.ResponseWriter, request *http.Request, key string) *string {
	queryStr := request.URL.Query().Get(key)
	if queryStr == "" {
		log.Println("query str empty.")
		*writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte(err.Error()))
	}

	return &queryStr
}
