package defect_features

import (
	"gaia/internal/jet/postgres/public/model"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"log"
)

type DefectCommandFeatureHandler struct {
	db provider.IDBProvider
}

func NewDefectCommandFeature(db provider.IDBProvider) *DefectCommandFeatureHandler {
	return &DefectCommandFeatureHandler{db: db}
}

func (commandHandler *DefectCommandFeatureHandler) InsertDefect(defect *model.Defects) error {
	insertStmt := table.Defects.INSERT(table.Defects.MutableColumns).MODEL(defect)
	err := commandHandler.db.Exec(insertStmt)
	if err != nil {
		log.Println("Unable to insert defects", err)
		return err
	}
	return nil
}
