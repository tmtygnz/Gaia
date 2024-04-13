package feature_handler

import (
	"gaia/internal/entities"
	defect_features "gaia/internal/features/defect_features"
	"gaia/internal/jet/postgres/public/model"
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

type DefectCommandFeature interface {
	InsertDefect(defect *entities.DDefects) error
}

type DefectRestImpl struct {
	defectQueryFeatureHandler   *defect_features.DefectQueryFeatureHandler
	defectCommandFeatureHandler *defect_features.DefectCommandFeatureHandler
}

func NewDefectRestHandler(defectFeatHandler *defect_features.DefectQueryFeatureHandler, defectCommandHandler *defect_features.DefectCommandFeatureHandler) {
	handler := &DefectRestImpl{
		defectQueryFeatureHandler:   defectFeatHandler,
		defectCommandFeatureHandler: defectCommandHandler,
	}
	http.HandleFunc("/defect", handler.FetchAllDefectHandler)
	http.HandleFunc("/defect/by/{id}", handler.FetchDefectByIdHandler)
	http.HandleFunc("/defect/search", handler.FetchFullTextSearchDefectHandler)

	log.Println("Defect rest handler created")
}

func (restHandler *DefectRestImpl) FetchAllDefectHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		allDefects, err := restHandler.defectQueryFeatureHandler.FetchAllDefects()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &allDefects, "application/json")
	}
}

func (restHandler *DefectRestImpl) FetchDefectByIdHandler(writer http.ResponseWriter, request *http.Request) {
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

		defect, err := restHandler.defectQueryFeatureHandler.FetchDefectById(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &defect, "application/json")
	}
}

func (restHandler *DefectRestImpl) FetchFullTextSearchDefectHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		queryStr := utils.GetRequestQuery(writer, request, "queryStr")
		if queryStr == nil {
			return
		}

		defects, err := restHandler.defectQueryFeatureHandler.FullTextSearchDefects(*queryStr)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &defects, "application/json")
	}
}

func (restHandler *DefectRestImpl) InsertDefectHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		defect := new(model.Defects)
		err := utils.ReadBody(writer, request, defect)
		if err != nil {
			return
		}

		err = restHandler.defectCommandFeatureHandler.InsertDefect(defect)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &defect, "application/json")
	}
}
