package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ProjectID    string `yaml:"project_id"`
	Zone         string `yaml:"zone"`
	InstanceName string `yaml:"instance_name"`
	MachineType  string `yaml:"machine_type"`
	DiskSizeGb   int64  `yaml:"disk_size_gb"`
	Image        string `yaml:"image"`
	NetworkName  string `yaml:"network_name"`
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
