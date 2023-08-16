package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DBUsername string `yaml:"dbUsername"`
	DBPassword string `yaml:"dbPassword"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	DBName     string `yaml:"dbName"`
}

func getDSN() string {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var dbConfig Config

	if err := yaml.Unmarshal(configFile, &dbConfig); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.Host, dbConfig.DBName)
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", getDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %s", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging DB: %s", err)
	}

	return db, nil
}
