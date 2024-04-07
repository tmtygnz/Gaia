package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/defect_type_features"
	"gaia/utils"
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

		utils.Send(writer, &allDefectType, "application/json")
	}
}

func (restHandler *DefectTypeRestHandler) FetchDefectTypeByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := utils.GetRequestQuery(writer, request, "id")
		if idStr == nil {
			return
		}

		id, err := strconv.Atoi(*idStr)
		if err != nil {
			http.Error(writer, "Invalid id type", http.StatusBadRequest)
			return
		}
		defectType := restHandler.defectFeatureHandler.FetchDefectTypeById(int64(id))

		utils.Send(writer, &defectType, "application/json")
	}
}
