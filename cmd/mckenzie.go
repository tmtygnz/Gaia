package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	defect_features "mckenzie/interal/features/defect-features"
	inline_rejection "mckenzie/interal/inline-rejection"
	"mckenzie/provider"
)

func startup() {
	log.Println("Loading .env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	startup()
	databaseInstance := provider.NewDatabase()
	defectFeatureHandler := defect_features.NewDefectFeatureHandler(databaseInstance)

	inline_rejection.IRRestInterface(defectFeatureHandler)
}
