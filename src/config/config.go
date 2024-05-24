package config

import (
	"encoding/json"
	"io"
	"log"
)

const (
	TypeJSON = "json"
)

type Config struct {
	Env                string `json:"env"`
	Port               string `json:"port"`
	AppName            string `json:"app_name"`
	DatabaseURL        string `json:"database_url"`
	MaxDBConn          int    `json:"max_db_conn"`
	TokenSize          int    `json:"token_size"`
	AccessTokenExpiry  int    `json:"access_token_expiry"`
	RefreshTokenExpiry int    `json:"refresh_token_expiry"`
}

var Conf *Config

func Parse(fType string, r io.Reader) error {
	switch fType {
	case TypeJSON:
		return ParseJson(r)
	}
	return nil
}

func ParseJson(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		log.Println("Unable to parse json config file.Err:", err)
		return err
	}
	Conf = &Config{}
	return json.Unmarshal(data, Conf)
}
