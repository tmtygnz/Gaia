package feature_handler

import (
	"encoding/json"
	"log"
	"mckenzie/interal/entities"
	defect_features "mckenzie/interal/features/defect-features"
	"net/http"
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
}

func (restHandler *DefectRestHandler) FetchAllDefectHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		allDefects := restHandler.defectFeatureHandler.FetchAllDefects()

		allDefectsBytes, err := json.Marshal(allDefects)
		if err != nil {
			log.Panicln("Error marshaling all defects", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.Write(allDefectsBytes)
	}
}
