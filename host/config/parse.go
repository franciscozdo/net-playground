package config

import (
  "os"
  "log"

  "gopkg.in/yaml.v2"
)

type ConfYaml struct {
  Name string
  Port int
  Connected, Services []string
}

func parseYaml(file string) (*ConfYaml, error) {

  dat, err := os.ReadFile(file)

  if err != nil {
    return nil, err
  }

  data := ConfYaml{}
  err = yaml.Unmarshal(dat, &data)
  if err != nil {
    return nil, err
  }

  log.Println(data)
  return &data, nil
}
