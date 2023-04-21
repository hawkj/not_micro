package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type Redis struct {
	Addr string `yaml:"addr"`
}

type Kafka struct {
	Host string `yaml:"host"`
}

type Config struct {
	MyTodoDb MySQLConfig `yaml:"my-todo-db"`
	Redis    Redis       `yaml:"redis"`
	Kafka    Kafka       `yaml:"kafka"`
}

func GetConfig(configFile string) *Config {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(content, &config); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}
	return &config
}

type Server struct {
	Name string
	Addr string
}

var Servers = map[string]Server{
	"gpt": {Name: "gpt", Addr: ":8901"},
}

func GetServerInfo(serverName string) *Server {
	server, ok := Servers[serverName]
	if !ok {
		return nil
	}
	return &server
}
