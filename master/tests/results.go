package tests

import (
  "net"
  "playground/master/config"
)

type TestResult struct {
  Addr net.IP
  Reachable *[]net.IP
  hasNames bool
  names []string
}

func parseResults(addr net.IP, reachable []string) TestResult {
  r := make([]net.IP, len(reachable))
  for i, a := range reachable {
    r[i] = net.ParseIP(a)
  }
  return TestResult{Addr: addr, Reachable: &r, hasNames: false}
}

func (tr TestResult) String() string {
  res := tr.Addr.String() + " can reach: "
  for i, a := range *tr.Reachable {
    res += a.String()
    if tr.hasNames {
      res += "(" + tr.names[i] + ")"
    }
    res += " "
  }
  return res
}

func (tr *TestResult) GetNames(conf *config.Config) {
  tr.names = make([]string, len(*tr.Reachable))

  for i, a := range *tr.Reachable {
    tr.names[i], _ = conf.GetName(a)
  }

  tr.hasNames = true
}
