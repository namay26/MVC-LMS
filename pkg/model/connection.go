package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

func getData(dbPath ...bool) string {
	var path string = "db.yaml"
	if dbPath != nil {
		path = "../../db.yaml"
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var databaseInfo struct {
		DB_USERNAME string `yaml:"DB_USERNAME"`
		DB_PASSWORD string `yaml:"DB_PASSWORD"`
		DB_HOST     string `yaml:"DB_HOST"`
		DB_NAME     string `yaml:"DB_NAME"`
	}

	if err := yaml.Unmarshal(file, &databaseInfo); err != nil {
		log.Fatal(err)
	}

	username := databaseInfo.DB_USERNAME
	password := databaseInfo.DB_PASSWORD
	hostname := databaseInfo.DB_HOST
	dbName := databaseInfo.DB_NAME

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)

}

func Connect(path ...bool) (*sql.DB, error) {
	db, err := sql.Open("mysql", getData(path[0]))
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	return db, err
}
