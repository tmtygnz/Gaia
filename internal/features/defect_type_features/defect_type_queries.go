package defect_type_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type DefectTypeQueryFeatureImpl struct {
	db provider.IDBProvider
}

func NewDefectTypeQueryFeature(db provider.IDBProvider) *DefectTypeQueryFeatureImpl {
	return &DefectTypeQueryFeatureImpl{db: db}
}

/*
FetchAllProducts fetches all products in the database
*/
func (handler *DefectTypeQueryFeatureImpl) FetchAllDefectType() (*[]entities.DDefectType, error) {
	var tfo = new([]entities.DDefectType)
	stmt := table.DefectsType.SELECT(table.DefectsType.AllColumns).FROM(table.DefectsType)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching all defect types", err)
		return nil, err
	}
	return tfo, nil
}

/*
FetchAllProducts fetches all products in the database
*/
func (handler *DefectTypeQueryFeatureImpl) FetchDefectTypeById(id int64) (*entities.DDefectType, error) {
	var tfo = new(entities.DDefectType)
	stmt := table.DefectsType.SELECT(table.DefectsType.AllColumns).FROM(table.DefectsType).
		WHERE(table.DefectsType.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching defect type by id", err)
		return nil, err
	}
	return tfo, nil
}
