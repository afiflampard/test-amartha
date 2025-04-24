package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //import postgres
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//DB ...

var db *gorm.DB

// Init ...
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	viperUser := os.Getenv("DB_USER")
	viperPassword := os.Getenv("DB_PASS")
	viperDb := os.Getenv("DB_NAME")
	viperHost := os.Getenv("DB_HOST")
	viperPort := os.Getenv("DB_PORT")

	prosgretConname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDb, viperPassword)

	db, err = ConnectDB(prosgretConname)
	if err != nil {
		log.Fatal(err)
	}

}

// ConnectDB ...
func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return db, nil
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}
