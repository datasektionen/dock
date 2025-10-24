package dao

import (
	"os"

	"github.com/datasektionen/dock/pkg/config"
	"gopkg.in/yaml.v3"
)

type Dao struct {
	cfg *config.Config
	Db  db
}

type db struct {
	Rfinger Rfinger `yaml:"rfinger"`
	// Ston    []User    `yaml:"ston"`
	// Spam    Hive      `yaml:"spam"`
}

type Rfinger struct {
	Pictures map[string]Picture `yaml:"pictures"`
	Default  string             `yaml:"default"`
}

type Picture struct {
	Regular string `yaml:"regular"`
	Small   string `yaml:"small"`
}

func New(cfg *config.Config) *Dao {
	file, err := os.ReadFile(cfg.ConfigFile)
	if err != nil {
		panic(err)
	}
	var db db
	err = yaml.Unmarshal(file, &db)
	if err != nil {
		panic(err)
	}

	return &Dao{
		Db: db,
	}
}

