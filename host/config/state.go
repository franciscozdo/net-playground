package config

import (
  "log"
  "net"
  "errors"
)

type Config struct {
  Name string
  Connected []net.IP
  Services []string
  Port int
}

func Init(file string) (*Config, error) {
  data, err := parseYaml(file)
  if err != nil {
    return nil, err
  }

  var conf Config
  conf.Name = data.Name
  conf.Services = data.Services
  conf.Port = data.Port

  conf.Connected = make([]net.IP, len(data.Connected))
  for i, addr := range data.Connected {
    host := net.ParseIP(addr)
    if host == nil {
      return nil, errors. New("unknown format of address in config")
    }
    conf.Connected[i] = host
    log.Printf("new host: %s", host.String())
  }
  return &conf, nil
}

func (c *Config) String() string {
  str := c.Name
  str += " connected ["
  for _, a := range c.Connected {
    str += a.String() + " "
  }
  str += "], services ["
  for _, s := range c.Services {
    str += s + " "
  }
  str += "]"
  return str
}
