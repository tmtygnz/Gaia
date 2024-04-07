package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/area_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type AreaQueries interface {
	FetchAllAreas() *[]entities.DArea
	FetchAreaById(id int64) *entities.DArea
}

type AreaRestHandler struct {
	areaQueries *area_features.AreaQueryHandler
}

func NewAreaRestHandler(areaQueryHandler *area_features.AreaQueryHandler) {
	handler := &AreaRestHandler{
		areaQueryHandler,
	}
	http.HandleFunc("/area", handler.FetchAllAreaHandler)
	http.HandleFunc("/area/by/", handler.FetchAreaByIdHandler)

	log.Println("Area rest handler created")
}

func (restHandler *AreaRestHandler) FetchAllAreaHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		areas := restHandler.areaQueries.FetchAllAreas()

		utils.Send(writer, &areas, "application/json")
	}
}

func (restHandler *AreaRestHandler) FetchAreaByIdHandler(writer http.ResponseWriter, request *http.Request) {
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
		area := restHandler.areaQueries.FetchAreaById(int64(id))

		utils.Send(writer, &area, "application/json")
	}
}
