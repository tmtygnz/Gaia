package feature_handler

import (
	"gaia/internal/entities"
	"gaia/internal/features/product_features"
	"gaia/utils"
	"log"
	"net/http"
	"strconv"
)

type ProductFeature interface {
	FetchAllProducts() *[]entities.DProduct
	FetchProductById(id int64) *entities.DProduct
}

type ProductRestHandler struct {
	productFeatureHandler *product_features.ProductFeatureHandler
}

func NewProductRestHandler(productFeatHandler *product_features.ProductFeatureHandler) {
	handler := &ProductRestHandler{
		productFeatureHandler: productFeatHandler,
	}
	http.HandleFunc("/product", handler.FetchAllProductHandler)
	http.HandleFunc("/product/by/", handler.FetchProductByIdHandler)

	log.Println("Product rest handler created")
}

func (restHandler *ProductRestHandler) FetchAllProductHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		products, err := restHandler.productFeatureHandler.FetchAllProducts()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		utils.Send(writer, &products, "application/json")
	}
}

func (restHandler *ProductRestHandler) FetchProductByIdHandler(writer http.ResponseWriter, request *http.Request) {
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

		product, err := restHandler.productFeatureHandler.FetchProductById(int64(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		utils.Send(writer, &product, "application/json")
	}
}
