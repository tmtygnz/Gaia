package packaging_features

import (
	"github.com/go-jet/jet/v2/postgres"
	"log"
	"mackenzie/internal/entities"
	"mackenzie/internal/jet/postgres/public/table"
	"mackenzie/provider"
)

type PackagingQueryHandler struct {
	db *provider.DBProvider
}

func NewPackagingTypeQueriesHandler(db *provider.DBProvider) *PackagingQueryHandler {
	return &PackagingQueryHandler{db: db}
}

func (handler *PackagingQueryHandler) FetchAllPackaging() *[]entities.DPackaging {
	var tfo = new([]entities.DPackaging)
	stmt := table.PackagingType.SELECT(table.PackagingType.AllColumns).FROM(table.PackagingType)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured whe fetching data")
		return nil
	}
	return tfo
}

func (handler *PackagingQueryHandler) FetchPackagingById(id int64) *entities.DPackaging {
	var tfo = new(entities.DPackaging)
	stmt := table.PackagingType.SELECT(table.PackagingType.AllColumns).FROM(table.PackagingType).
		WHERE(table.PackagingType.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured when fetching defect type by id", err)
		return nil
	}
	return tfo
}
