package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Env int

const (
	_ Env = iota
	Dev
	Stage
	Prod
)

var envFromString = map[string]Env{
	"dev":         Dev,
	"development": Dev,
	"stage":       Stage,
	"staging":     Stage,
	"prod":        Prod,
	"production":  Prod,
}

func (e Env) String() string {
	switch e {
	case Dev:
		return "dev"
	case Stage:
		return "stage"
	case Prod:
		return "prod"
	default:
		return "dev"
	}
}

type Config struct {
	Database struct {
		User     string `json:"user"`
		Password string `json:"password,omitempty"`
		Host     string `json:"host,omitempty"`
		Port     string `json:"port,omitempty"`
		Database string `json:"database"`
	} `json:"database"`
	OAuth struct {
		Secret string `json:"secret"`
		ID     string `json:"id"`
	} `json:"oauth"`
	Environment string
}

func Load() Config {
	env := envFromString[os.Getenv("GO_ENV")]
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	cfg_file := filepath.Join(dir, fmt.Sprintf("%v.json", env))
	data, err := ioutil.ReadFile(cfg_file)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	fmt.Printf("%v", cfg)
	return cfg
}
