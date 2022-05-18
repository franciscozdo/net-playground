package main

import (
  "log"
  "flag"
  "sync"

  "playground/host/server"
  "playground/host/config"
  "playground/host/tests"
)

func ParseArgs() (string) {
  conf := flag.String("c", "config.yaml", "configuration file")

  flag.Parse()

  return *conf
}

func main() {
  configFile := ParseArgs()

  conf, err := config.Init(configFile)
  if err != nil {
    log.Fatalf("Error during config initialization (%v)", err)
  }
  log.Print(conf)

  /* I want to buffer only first channels because they may cause sleep of server routine */
  newTest := make(chan tests.TestEntry, 128)
  testReq := make(chan tests.TestEntry, 128)
  testRes := make(chan tests.TestEntry)
  dataReq := make(chan tests.TestEntry)
  dataRes := make(chan tests.TestEntry)

  dChannels := tests.DataRoutineChannels{ NewTest: newTest, DataReq: dataReq, DataRes: dataRes, TestRes: testRes }
  tChannels := tests.TestRoutineChannels{ TestReq: testReq, TestRes: testRes }
  sChannels := server.ServerRoutineChannels{ DataReq: dataReq, DataRes: dataRes, TestReq: testReq, NewTest: newTest }

  var wg sync.WaitGroup
  wg.Add(2 + 10)
  go tests.DataRoutine(dChannels, wg)
  for i := 0; i < 10; i++ {
    go tests.TestRoutine(tChannels, wg, conf)
  }
  go server.ServerRoutine(sChannels, wg, conf)

  wg.Wait()
}
