package packaging_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type PackagingQueryHandler struct {
	db *provider.DBProvider
}

func NewPackagingTypeQueriesHandler(db *provider.DBProvider) *PackagingQueryHandler {
	return &PackagingQueryHandler{db: db}
}

func (handler *PackagingQueryHandler) FetchAllPackaging() (*[]entities.DPackaging, error) {
	var tfo = new([]entities.DPackaging)
	stmt := table.PackagingType.SELECT(table.PackagingType.AllColumns).FROM(table.PackagingType)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured whe fetching data")
		return nil, err
	}
	return tfo, nil
}

func (handler *PackagingQueryHandler) FetchPackagingById(id int64) (*entities.DPackaging, error) {
	var tfo = new(entities.DPackaging)
	stmt := table.PackagingType.SELECT(table.PackagingType.AllColumns).FROM(table.PackagingType).
		WHERE(table.PackagingType.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured when fetching defect type by id", err)
		return nil, err
	}
	return tfo, nil
}
