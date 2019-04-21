package RBAC

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)


type Configuration struct {
    DBWriter        DBWriter        `yaml:"db_writer"`
    DBReader        DBReader        `yaml:"db_reader"`
}

type DBWriter struct {
    Host            string      `yaml:"host"`
    Name            string      `yaml:"name"`
    Username        string      `yaml:"username"`
    Password        string      `yaml:"password"`
    Port            string      `yaml:"port"`
}

type DBReader struct {
    Host            string      `yaml:"host"`
    Name            string      `yaml:"name"`
    Username        string      `yaml:"username"`
    Password        string      `yaml:"password"`
    Port            string      `yaml:"port"`
}

var (
    Config Configuration
)

func LoadConfiguration() {
    configFile, err := ioutil.ReadFile("conf/dev.yaml")

    if err != nil {
        fmt.Println("Unable to open file")
    }

    err = yaml.Unmarshal(configFile, &Config)

    if err != nil {
        fmt.Println("Unable to parse config")
    }

    return
}
