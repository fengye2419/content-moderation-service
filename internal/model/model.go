package model

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

// InitializeEngine initializes the xorm engine.
func InitializeEngine(configPath string) {
	cfg, err := ini.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	dbUser := cfg.Section("database").Key("user").String()
	dbPassword := cfg.Section("database").Key("password").String()
	dbHost := cfg.Section("database").Key("host").String()
	dbName := cfg.Section("database").Key("name").String()
	dbCharset := cfg.Section("database").Key("charset").String()

	dataSourceName := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=" + dbCharset

	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to create engine: %v", err)
	}
	log.Println("Database connection established successfully")

	// Set the engine's mapper to GonicMapper
	engine.SetMapper(names.GonicMapper{})

	// Enable SQL log output
	engine.ShowSQL(true)
}
