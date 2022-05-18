package config

import (
  "os"

  "gopkg.in/yaml.v2"
)

type HostYaml struct {
  Address string
  Port int
  Connected []string
  Services []string
  Connections int
}

func parseYaml(file string) (map[string]HostYaml, error) {

  dat, err := os.ReadFile(file)

  if err != nil {
    return nil, err
  }

  data := make(map[string]HostYaml)
  err = yaml.Unmarshal(dat, &data)
  if err != nil {
    return nil, err
  }

  return data, nil
}
