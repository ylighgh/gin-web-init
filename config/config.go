package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

var G *Config

type Config struct {
	SonarQube SonarQube `toml:"sonarqube"`
}

type SonarQube struct {
	AbsoluteExecPath string `toml:"absolute_exec_path"`

	Key string `toml:"key"`

	Host string `toml:"host"`
}

func ParseConfig(file string) (*Config, error) {
	fi, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	_ = fi

	cnt, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cnf := &Config{}

	if err = toml.Unmarshal(cnt, cnf); err != nil {
		return nil, err
	}

	return cnf, nil
}

func InitConfig() {
	var def = os.Getenv("CONFIG")
	if def == "" {
		def = "./config.toml"
	}

	cnf, err := ParseConfig(def)
	if err != nil {
		fmt.Printf("could not parse config from file %s, error = %s\n", def, err.Error())
		return
	}

	G = cnf
}
