package utils

import (
	"log"
	"net/http"
	"html/template"
)

func RenderPage(name string, pageLoc string, writer http.ResponseWriter, data interface{}) error {
	t := template.New(name)
	temp, err := t.ParseGlob(pageLoc)
	if err != nil {
		log.Println(err)
	}
	return temp.Execute(writer, data)
}
