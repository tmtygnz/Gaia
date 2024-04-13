package product_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type ProductFeatureHandler struct {
	db provider.IDBProvider
}

func NewProductFeatureHandler(db provider.IDBProvider) *ProductFeatureHandler {
	return &ProductFeatureHandler{db: db}
}

/*
FetchAllProducts fetches all products in the database
*/
func (handler *ProductFeatureHandler) FetchAllProducts() (*[]entities.DProduct, error) {
	var tfo = new([]entities.DProduct)
	stmt := table.Products.SELECT(table.Products.AllColumns).FROM(table.Products)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching all products", err)
		return nil, err
	}
	return tfo, nil
}

/*
FetchProductById fetches product with the same id as the parameter
*/
func (handler *ProductFeatureHandler) FetchProductById(id int64) (*entities.DProduct, error) {
	tfo := new(entities.DProduct)
	stmt := table.Products.SELECT(table.Products.AllColumns).FROM(table.Products).
		WHERE(table.Products.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching product by id", err)
		return nil, err
	}
	return tfo, nil
}
