package packaging_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type PackagingQueryFeatureImpl struct {
	db *provider.DBProvider
}

func NewPackagingTypeQuery(db *provider.DBProvider) *PackagingQueryFeatureImpl {
	return &PackagingQueryFeatureImpl{db: db}
}

func (handler *PackagingQueryFeatureImpl) FetchAllPackaging() (*[]entities.DPackaging, error) {
	var tfo = new([]entities.DPackaging)
	stmt := table.PackagingType.SELECT(table.PackagingType.AllColumns).FROM(table.PackagingType)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured whe fetching data")
		return nil, err
	}
	return tfo, nil
}

func (handler *PackagingQueryFeatureImpl) FetchPackagingById(id int64) (*entities.DPackaging, error) {
	var tfo = new(entities.DPackaging)
	stmt := table.PackagingType.SELECT(table.PackagingType.AllColumns).FROM(table.PackagingType).
		WHERE(table.PackagingType.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured when fetching defect type by id", err)
		return nil, err
	}
	return tfo, nil
}
