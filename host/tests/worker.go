package tests

import (
  "net"
  "log"
  "sync"

  msg "playground/host/messages"
  "playground/host/config"
)

type TestEntry struct {
  Type msg.Test
  Id int
  Done bool
  Data *[]string
}

type TestRoutineChannels struct {
  TestReq <-chan TestEntry
  TestRes chan<- TestEntry
}

func TestRoutine(channels TestRoutineChannels, wg sync.WaitGroup, conf *config.Config) {
  defer wg.Done()
  for {
    req := <-channels.TestReq
    log.Printf("Running test %d %s", req.Id, req.Type.String())
    res := RunTest(req.Type, conf)
    channels.TestRes <- TestEntry{Id: req.Id, Done: true, Data: &res}
  }
}

type testFunType func(net.IP) error

func TestMany(hosts []net.IP, fn testFunType) ([]string) {
  reachable := make([]string, len(hosts))
  i := 0
  for _, h := range hosts {
    err := fn(h)
    log.Print(err)
    if err == nil {
      reachable[i] = h.String()
      i++
    }
  }
  return reachable[:i]
}

func RunTest(tp msg.Test, conf *config.Config) ([]string) {
  var fn testFunType
  switch tp {
  case msg.Test_ECHO:
    fn = EchoTest
  case msg.Test_HTTP:
    fn = HttpTest
  }
  return TestMany(conf.Connected, fn)
}

type DataRoutineChannels struct {
  NewTest, TestRes, DataReq <-chan TestEntry
  DataRes chan<- TestEntry
}

func DataRoutine(channels DataRoutineChannels, wg sync.WaitGroup) {
  defer wg.Done()
  storage := make(map[int]*TestEntry, 128)
  for {
    select {
    case nt := <-channels.NewTest:
      if _, ok := storage[nt.Id]; ok { continue }
      entry := nt
      storage[nt.Id] = &entry
    case tr := <-channels.TestRes:
      if _, ok := storage[tr.Id]; !ok {
        log.Printf("Result from unknown test %d", tr.Id)
        continue
      }
      storage[tr.Id].Done = true
      storage[tr.Id].Data = tr.Data
    case dr := <-channels.DataReq:
      if _, ok := storage[dr.Id]; !ok {
        log.Printf("Data request for unknown test %d", dr.Id)
        continue
      }
      channels.DataRes <- *storage[dr.Id]
    }
  }
}
