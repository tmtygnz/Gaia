package defect_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type DefectQueryFeatureHandler struct {
	db provider.IDBProvider
}

func NewDefectQueryFeatureHandler(db provider.IDBProvider) *DefectQueryFeatureHandler {
	return &DefectQueryFeatureHandler{db: db}
}

func (handler *DefectQueryFeatureHandler) FetchAllDefects() (*[]entities.DDefects, error) {
	var tfo = new([]entities.DDefects)
	stmt := table.Defects.SELECT(table.Defects.AllColumns).FROM(table.Defects)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching all defects from the database", err)
		return nil, err
	}
	return tfo, nil
}

func (handler *DefectQueryFeatureHandler) FetchDefectById(id int64) (*entities.DDefects, error) {
	var tfo = new(entities.DDefects)
	stmt := table.Defects.SELECT(table.Defects.AllColumns).FROM(table.Defects).
		WHERE(table.Defects.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred while fetching defect by id", err)
		return nil, err
	}
	return tfo, nil
}

func (handler *DefectQueryFeatureHandler) FullTextSearchDefects(query string) (*[]entities.DDefects, error) {
	var tfo = new([]entities.DDefects)
	stmt := table.Defects.SELECT(table.Defects.AllColumns).FROM(table.Defects).WHERE(table.Defects.DefectDescription.EQ(postgres.String(query)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured while fetching defect via query", err)
		return nil, err
	}
	return tfo, nil
}
