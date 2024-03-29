package feature_handler

import (
	"encoding/json"
	"log"
	"mackenzie/internal/entities"
	defect_features "mackenzie/internal/features/defect_features"
	"net/http"
	"strconv"
)

type DefectFeature interface {
	FetchAllDefects() *[]entities.DDefects
	FetchDefectById(id int64) *entities.DDefects
}

type DefectRestHandler struct {
	defectFeatureHandler *defect_features.DefectFeatureHandler
}

func NewDefectRestHandler(defectFeatHandler *defect_features.DefectFeatureHandler) {
	handler := &DefectRestHandler{
		defectFeatureHandler: defectFeatHandler,
	}
	http.HandleFunc("/defect", handler.FetchAllDefectHandler)
	http.HandleFunc("/defect/by/{id}", handler.FetchDefectByIdHandler)
}

func (restHandler *DefectRestHandler) FetchAllDefectHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		allDefects := restHandler.defectFeatureHandler.FetchAllDefects()

		allDefectsBytes, err := json.Marshal(allDefects)
		if err != nil {
			log.Println("Error marshaling all defects", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.Write(allDefectsBytes)
	}
}

func (restHandler *DefectRestHandler) FetchDefectByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := request.URL.Query().Get("id")
		if idStr == "" {
			writer.WriteHeader(http.StatusUnprocessableEntity)
			writer.Write([]byte("Missing id parameter"))
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		defect := restHandler.defectFeatureHandler.FetchDefectById(int64(id))

		defectBytes, err := json.Marshal(defect)
		if err != nil {
			log.Println("Error marshaling defect", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.Write(defectBytes)
	}
}
