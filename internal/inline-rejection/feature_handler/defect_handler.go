package feature_handler

import (
	"gaia/internal/entities"
	defect_features "gaia/internal/features/defect_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type DefectQueryFeature interface {
	FetchAllDefects() (*[]entities.DDefects, error)
	FetchDefectById(id int64) (*entities.DDefects, error)
	FullTextSearchDefects(query string) (*[]entities.DDefects, error)
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
		allDefects, err := restHandler.defectFeatureHandler.FetchAllDefects()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

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

		defect, err := restHandler.defectFeatureHandler.FetchDefectById(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

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

		defects, err := restHandler.defectFeatureHandler.FullTextSearchDefects(*queryStr)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &defects, "application/json")
	}
}
