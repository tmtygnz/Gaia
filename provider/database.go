package provider

import (
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	"log"
	"os"
)

type IDBProvider interface {
	Query(query postgres.SelectStatement, tfo interface{}) error
	Exec(query postgres.Statement, tfo interface{}) error
}

type DBProvider struct {
	database *sql.DB
}

func NewDatabase() *DBProvider {
	db, err := sql.Open("postgres", os.Getenv("DB_STR"))
	if err != nil {
		log.Fatal("Cannot connect to postgres", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Didn't connect to the database", err)
	}

	return &DBProvider{
		database: db,
	}
}

/*
Query the given select statement.
Use these for functions that returns rows.

Usage Example:
stmt := Link.Select(Link.AllColumns).From(Link)
var tfo []model.Link

database.Query(stmt, tfo)
*/
func (prov *DBProvider) Query(query postgres.SelectStatement, tfo interface{}) error {
	err := query.Query(prov.database, tfo)
	if err != nil {
		log.Panicln("Query error", err)
		return err
	}
	return nil
}

/*
Exec executes the given statement.
Use these for functions that doesn't return rows
*/
func (prov *DBProvider) Exec(query postgres.Statement, tfo interface{}) error {
	_, err := query.Exec(prov.database)
	if err != nil {
		log.Panicln("Exec error", err)
		return err
	}
	return nil
}
