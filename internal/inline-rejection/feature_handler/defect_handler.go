package feature_handler

import (
	"gaia/internal/entities"
	defect_features "gaia/internal/features/defect_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type DefectFeature interface {
	FetchAllDefects() *[]entities.DDefects
	FetchDefectById(id int64) *entities.DDefects
}

type DefectRestHandler struct {
	defectFeatureHandler *defect_features.DefectQueryFeatureHandler
}

func NewDefectRestHandler(defectFeatHandler *defect_features.DefectQueryFeatureHandler) {
	handler := &DefectRestHandler{
		defectFeatureHandler: defectFeatHandler,
	}
	http.HandleFunc("/defect", handler.FetchAllDefectHandler)
	http.HandleFunc("/defect/by/{id}", handler.FetchDefectByIdHandler)
	http.HandleFunc("/defect/search", handler.FetchFullTextSearchDefectHandler)

	log.Println("Defect rest handler created")
}

func (restHandler *DefectRestHandler) FetchAllDefectHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		allDefects := restHandler.defectFeatureHandler.FetchAllDefects()

		utils.Send(writer, &allDefects, "application/json")
	}
}

func (restHandler *DefectRestHandler) FetchDefectByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := utils.GetRequestQuery(writer, request, "id")
		if idStr == nil {
			return
		}

		id, err := strconv.Atoi(*idStr)
		if err != nil {
			http.Error(writer, "Invalid id type", http.StatusInternalServerError)
		}
		defect := restHandler.defectFeatureHandler.FetchDefectById(int64(id))

		utils.Send(writer, &defect, "application/json")
	}
}

func (restHandler *DefectRestHandler) FetchFullTextSearchDefectHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		queryStr := utils.GetRequestQuery(writer, request, "queryStr")
		if queryStr == nil {
			return
		}

		defects := restHandler.defectFeatureHandler.FullTextSearchDefects(*queryStr)

		utils.Send(writer, &defects, "application/json")
	}
}
