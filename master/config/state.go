package config

import (
  "fmt"
  "net"
  . "playground/master/ferrors"
  msg "playground/master/messages"
)

type Host struct {
  Address net.IP
  Port int
  Name string
  Connected []net.IP
  ConnectedName []string
  Services []msg.Test
}

type Config struct {
  Hosts map[string]Host
  MaxConnections int
}

func str2ttype(s string) msg.Test {
  switch s {
  case "echo":
    return msg.Test_ECHO
  case "http":
    return msg.Test_HTTP
  default:
    return msg.Test_UNKNOWN
  }
}

func Init(file string) (*Config, error) {
  hosts, err := parseYaml(file)
  if err != nil {
    return nil, MakeError(ErrConfig, err.Error())
  }

  var conf Config

  conf.MaxConnections = hosts["master"].Connections
  conf.Hosts = make(map[string]Host)

  for name, data := range hosts {
    if name == "master" {
      continue
    }

    addr := net.ParseIP(data.Address)
    if addr == nil {
      return nil, MakeError(ErrConfig, "config: invalid ip address")
    }

    nConnections := len(data.Connected)
    connected := make([]net.IP, nConnections)
    for i := range data.Connected {
      /* connected hosts are given by name so we have to obtain their address */
      cn := data.Connected[i]
      ca := net.ParseIP(hosts[cn].Address)
      if ca == nil {
        return nil, MakeError(ErrConfig, "config: invalid ip address")
      }
      connected[i] = ca
    }

    services := make([]msg.Test, len(data.Services))
    for i, srv := range data.Services {
      services[i] = str2ttype(srv)
    }

    host := Host{Address: addr, Port: data.Port, Connected: connected, ConnectedName: data.Connected, Services: services}
    conf.Hosts[name] = host
  }

  return &conf, nil
}

func (conf *Config) Print() {
  fmt.Printf("Max connections: %d\n", conf.MaxConnections)
  fmt.Println("Known hosts:")

  for name, data := range conf.Hosts {
    fmt.Printf("  - %s (%s:%d) services: %s\n", name, data.Address, data.Port, data.Services)
    fmt.Print("    Connected to:")
    for i := range data.Connected {
      fmt.Printf(" (%s %s)", data.ConnectedName[i], data.Connected[i])
    }
    fmt.Println("")
  }
  fmt.Println("")
}

func (conf *Config) GetAddress(host string) (net.IP, int, error) {
  a, ok := conf.Hosts[host]
  if !ok {
    return nil, 0, MakeError(ErrConfig, fmt.Sprintf("host '%s' not found", host))
  }
  return a.Address, a.Port, nil
}

func (conf *Config) GetName(addr net.IP) (string, error) {
  for name, host := range conf.Hosts {
    if addr.Equal(host.Address) {
      return name, nil
    }
  }
  return "", MakeError(ErrConfig, "host not found")
}
