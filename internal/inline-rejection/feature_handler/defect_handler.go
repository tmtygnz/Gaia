package feature_handler

import (
	"gaia/internal/entities"
	defect_features "gaia/internal/features/defect_features"
	"gaia/utils"
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

		utils.Send(writer, &allDefects, "application/json")
	}
}

func (restHandler *DefectRestHandler) FetchDefectByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := utils.GetRequestQuery(writer, request, "id")

		id, err := strconv.Atoi(*idStr)
		if err != nil {
			http.Error(writer, "Invalid id type", http.StatusInternalServerError)
		}
		defect := restHandler.defectFeatureHandler.FetchDefectById(int64(id))

		utils.Send(writer, &defect, "application/json")
	}
}
