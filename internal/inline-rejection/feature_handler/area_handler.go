package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/area_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type AreaQueryFeature interface {
	FetchAllAreas() (*[]entities.DArea, error)
	FetchAreaById(id int64) (*entities.DArea, error)
}

type AreaRestImpl struct {
	areaQueries *area_features.AreaQueryFeatureImpl
}

func NewAreaRestHandler(areaQueryHandler *area_features.AreaQueryFeatureImpl) {
	handler := &AreaRestImpl{
		areaQueryHandler,
	}
	http.HandleFunc("/area", handler.FetchAllAreaHandler)
	http.HandleFunc("/area/by/", handler.FetchAreaByIdHandler)

	log.Println("Area rest handler created")
}

func (restHandler *AreaRestImpl) FetchAllAreaHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		areas, err := restHandler.areaQueries.FetchAllAreas()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &areas, "application/json")
	}
}

func (restHandler *AreaRestImpl) FetchAreaByIdHandler(writer http.ResponseWriter, request *http.Request) {
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

		area, err := restHandler.areaQueries.FetchAreaById(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &area, "application/json")
	}
}
