package config

import "github.com/kelseyhightower/envconfig"

type DB struct {
	Port    string `split_words:"true" default:"8080"`
	LogFile string `split_words:"true" default:"webapp.log"`
	Driver  string `split_words:"true" default:"sqlite3"`
	Name    string `split_words:"true" default:"webapp.sql"`
}

func NewDB() (DB, error) {
	conf := DB{}
	if err := envconfig.Process("DB", &conf); err != nil {
		return conf, err
	}
	return conf, nil
}
