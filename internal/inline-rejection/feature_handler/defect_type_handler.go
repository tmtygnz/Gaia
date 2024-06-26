package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/defect_type_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type DefectTypeQueryFeature interface {
	FetchAllDefectType() (*[]entities.DDefectType, error)
	FetchDefectTypeById(id int64) (*entities.DDefectType, error)
}

type DefectTypeRestImpl struct {
	defectFeatureHandler *defect_type_features.DefectTypeQueryFeatureImpl
}

func NewDefectTypeRestHandler(defectFeatureHandler *defect_type_features.DefectTypeQueryFeatureImpl) {
	handler := &DefectTypeRestImpl{
		defectFeatureHandler: defectFeatureHandler,
	}
	http.HandleFunc("/defectType", handler.FetchAllDefectTypeHandler)
	http.HandleFunc("/defectType/by/{id}", handler.FetchDefectTypeByIdHandler)

	log.Println("Defect type rest handler created")
}

func (restHandler *DefectTypeRestImpl) FetchAllDefectTypeHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		allDefectType, err := restHandler.defectFeatureHandler.FetchAllDefectType()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		utils.Send(writer, &allDefectType, "application/json")
	}
}

func (restHandler *DefectTypeRestImpl) FetchDefectTypeByIdHandler(writer http.ResponseWriter, request *http.Request) {
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

		defectType, err := restHandler.defectFeatureHandler.FetchDefectTypeById(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &defectType, "application/json")
	}
}
