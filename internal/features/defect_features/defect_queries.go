package defect_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type DefectFeatureHandler struct {
	db provider.IDBProvider
}

func NewDefectFeatureHandler(db provider.IDBProvider) *DefectFeatureHandler {
	return &DefectFeatureHandler{db: db}
}

func (handler *DefectFeatureHandler) FetchAllDefects() *[]entities.DDefects {
	var tfo = new([]entities.DDefects)
	stmt := table.Defects.SELECT(table.Defects.AllColumns).FROM(table.Defects)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching all defects from the database", err)
		return nil
	}
	return tfo
}

func (handler *DefectFeatureHandler) FetchDefectById(id int64) *entities.DDefects {
	var tfo = new(entities.DDefects)
	stmt := table.Defects.SELECT(table.Defects.AllColumns).FROM(table.Defects).
		WHERE(table.Defects.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred while fetching defect by id", err)
		return nil
	}
	return tfo
}

func (handler *DefectFeatureHandler) FullTextSearchDefects(query string) *[]entities.DDefects {
	var tfo = new([]entities.DDefects)
	stmt := table.Defects.SELECT(table.Defects.AllColumns).FROM(table.Defects).WHERE(table.Defects.DefectDescription.EQ(postgres.String(query)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occured while fetching defect via query", err)
		return nil
	}
	return tfo
}
