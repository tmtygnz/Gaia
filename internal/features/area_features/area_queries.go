package area_features

import (
	"gaia/internal/entities"
	"gaia/internal/jet/postgres/public/table"
	"gaia/provider"
	"github.com/go-jet/jet/v2/postgres"
	"log"
)

type AreaQueryHandler struct {
	db provider.IDBProvider
}

func NewAreaFeatureHandler(db provider.IDBProvider) *AreaQueryHandler {
	return &AreaQueryHandler{db: db}
}

func (handler *AreaQueryHandler) FetchAllAreas() *[]entities.DArea {
	var tfo = new([]entities.DArea)
	stmt := table.Places.SELECT(table.Places.AllColumns).FROM(table.Places)
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred when fetching all areas from the database", err)
		return nil
	}
	return tfo
}

func (handler *AreaQueryHandler) FetchAreaById(id int64) *entities.DArea {
	var tfo = new(entities.DArea)
	stmt := table.Places.SELECT(table.Places.AllColumns).FROM(table.Places).
		WHERE(table.Places.ID.EQ(postgres.Int64(id)))
	if err := handler.db.Query(stmt, tfo); err != nil {
		log.Println("An error occurred while fetching area by id", err)
		return nil
	}
	return tfo
}
