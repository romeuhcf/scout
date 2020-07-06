package main

import (
	"io/ioutil"

	"encoding/json"
)

// Config is the internal representation of the json that determines what
// the app listens to an enqueues
type Config struct {
	Redis RedisConfig `json:"redis"`
	AWS   AWSConfig   `json:"aws"`
	Queue QueueConfig `json:"queue"`
}

// RedisConfig is a nested config that contains the necessary parameters to
// connect to a redis instance and enqueue workers.
type RedisConfig struct {
	Host      string `json:"host"`
	Namespace string `json:"namespace"`
	Queue     string `json:"queue"`
}

// AWSConfig is a nested config that contains the necessary parameters to
// connect to AWS and read from SQS
type AWSConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Region    string `json:"region"`
}

// QueueConfig is a nested config that gives the SQS queue to listen on
// and a mapping of topics to workeers
type QueueConfig struct {
	Name   string            `json:"name"`
	Topics map[string]string `json:"topics"`
}

// ReadConfig reads from a file with the given name and returns a config or
// an error if the file was unable to be parsed. It does no error checking
// as far as required fields.
func ReadConfig(file string) (*Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	err = json.Unmarshal(data, config)
	return config, err
}
