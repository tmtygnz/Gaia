package feature_handler

import (
	"encoding/json"
	"gaia/internal/entities"
	"gaia/internal/features/defect_type_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type DefectTypeFeature interface {
	FetchAllDefectType() *[]entities.DDefectType
	FetchDefectTypeById(id int64) *entities.DDefectType
}

type DefectTypeRestHandler struct {
	defectFeatureHandler *defect_type_features.DefectTypeFeatureHandler
}

func NewDefectTypeRestHandler(defectFeatureHandler *defect_type_features.DefectTypeFeatureHandler) {
	handler := &DefectTypeRestHandler{
		defectFeatureHandler: defectFeatureHandler,
	}
	http.HandleFunc("/defectType", handler.FetchAllDefectTypeHandler)
	http.HandleFunc("/defectType/by/{id}", handler.FetchDefectTypeByIdHandler)
}

func (restHandler *DefectTypeRestHandler) FetchAllDefectTypeHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		allDefectType := restHandler.defectFeatureHandler.FetchAllDefectType()

		allDefectTypeBytes, err := json.Marshal(allDefectType)
		if err != nil {
			log.Println("Something went wrong when marshaling all defect types", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		if _, err := writer.Write(allDefectTypeBytes); err != nil {
			log.Println("Failed writing all defect type bytes to http writer", err)
		}
	}
}

func (restHandler *DefectTypeRestHandler) FetchDefectTypeByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := utils.GetRequestQuery(writer, request, "id")

		id, err := strconv.Atoi(*idStr)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		defectType := restHandler.defectFeatureHandler.FetchDefectTypeById(int64(id))

		defectTypeBytes, err := json.Marshal(defectType)
		if err != nil {
			log.Println("Can't marshal defectType")
			writer.WriteHeader(http.StatusInternalServerError)
		}

		writer.Write(defectTypeBytes)
	}
}
