package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/packaging_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type PackagingQueryFeature interface {
	FetchAllPackaging() *[]entities.DPackaging
	FetchPackagingById(id int64) *entities.DPackaging
}

type PackagingRestImpl struct {
	packagingQueryHandler *packaging_features.PackagingQueryFeatureImpl
}

func NewPackagingRestHandler(packagingQueryHandler *packaging_features.PackagingQueryFeatureImpl) {
	handler := &PackagingRestImpl{
		packagingQueryHandler,
	}
	http.HandleFunc("/packaging", handler.FetchAllPackagingHandler)
	http.HandleFunc("/packaging/by/", handler.FetchPackagingByIdHandler)

	log.Println("Packaging rest handler created")
}

func (restHandler *PackagingRestImpl) FetchAllPackagingHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		packaging, err := restHandler.packagingQueryHandler.FetchAllPackaging()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		utils.Send(writer, &packaging, "application/json")
	}
}

func (restHandler *PackagingRestImpl) FetchPackagingByIdHandler(writer http.ResponseWriter, request *http.Request) {
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

		packaging, err := restHandler.packagingQueryHandler.FetchPackagingById(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &packaging, "application/json")
	}
}
