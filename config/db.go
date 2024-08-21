package config

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"gopkg.in/yaml.v2"

	_ "github.com/lib/pq"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"db"`
}

// docker run --name postgres-schoolapp -e POSTGRES_DB=schoolApp -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123123 -p 5432:5432 -d postgres
var DB *sql.DB

func LoadConfig() (*Config, error) {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ConnectDB() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Connected to the database successfully")
}
