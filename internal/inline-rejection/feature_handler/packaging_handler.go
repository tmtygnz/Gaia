package feature_handler

import (
	"encoding/json"
	"gaia/internal/entities"
	"gaia/internal/features/packaging_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type PackagingQueries interface {
	FetchAllPackaging() *[]entities.DPackaging
	FetchPackagingById(id int64) *entities.DPackaging
}

type PackagingRestHandler struct {
	packagingQueryHandler *packaging_features.PackagingQueryHandler
}

func NewPackagingRestHandler(packagingQueryHandler *packaging_features.PackagingQueryHandler) {
	handler := &PackagingRestHandler{
		packagingQueryHandler,
	}
	http.HandleFunc("/packaging", handler.FetchAllPackagingHandler)
	http.HandleFunc("/packaging/by/", handler.FetchPackagingByIdHandler)
}

func (restHandler *PackagingRestHandler) FetchAllPackagingHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		allProducts := restHandler.packagingQueryHandler.FetchAllPackaging()

		allPackagingBytes, err := json.Marshal(allProducts)
		if err != nil {
			http.Error(writer, "Server is unable to marshal packaging to bytes", http.StatusInternalServerError)
			return
		}

		utils.Send(writer, &allPackagingBytes, "application/json")
	}
}

func (restHandler *PackagingRestHandler) FetchPackagingByIdHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		idStr := utils.GetRequestQuery(writer, request, "id")

		id, err := strconv.Atoi(*idStr)
		if err != nil {
			http.Error(writer, "Invalid id type", http.StatusBadRequest)
			return
		}
		packaging := restHandler.packagingQueryHandler.FetchPackagingById(int64(id))

		packagingBytes, err := json.Marshal(packaging)
		if err != nil {
			log.Println("Can't marshal packaging to bytes", err)
			http.Error(writer, "Server is unable to marshal packaging to bytes", http.StatusInternalServerError)
			return
		}

		utils.Send(writer, &packagingBytes, "application/json")
	}
}
