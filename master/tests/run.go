package tests

import (
  "log"
  "time"

  "playground/master/client"
  "playground/master/config"
  msg "playground/master/messages"
  . "playground/master/ferrors"
)

var nextId int32 = 0

func idGen() int32 {
  nextId++
  return nextId
}

func runOneTest(id int32, ttype msg.Test, c *client.Client) ([]string, error) {
  if ttype == msg.Test_UNKNOWN {
    return nil, MakeError(ErrTests, "unknown test to schedule")
  }

	err := c.RunTest(id, ttype)

  if err != nil {
    return nil, err
  }

  log.Printf("Test %d scheduled...\n", id)

  for {
    time.Sleep(3 * time.Second)
    done, resp, err := c.GetResult(id)

    if err != nil {
      return nil, err
    }

    if done {
      return resp, nil
    }
  }
}

func TestOne(ttype msg.Test, c *client.Client, conf *config.Config) (string, error) {
  id := idGen()

  reachable, err := runOneTest(id, ttype, c)
  if err != nil {
    return "", err
  }

  res := parseResults(c.Addr, reachable)
  res.GetNames(conf)

  return res.String(), nil
}

func testingRoutine(id int32, ttype msg.Test, c *client.Client, ch chan<- TestResult) {
  reachable, err := runOneTest(id, ttype, c)
  if err != nil {
    log.Printf("routine for client %s: %v", c.Addr, err)
  }
  res := parseResults(c.Addr, reachable)
  ch <- res
}

func TestMany(ttype msg.Test, clState *client.State, conf *config.Config) (string, error) {
  ch := make(chan TestResult)
  id := idGen()

  for i := 0; i < clState.NConnections; i++ {
    c := clState.Connections[i]
    go testingRoutine(id, ttype, c, ch)
  }

  msg := ""
  for i := 0; i < clState.NConnections; i++ {
    res := <-ch
    res.GetNames(conf)
    msg += res.String() + "\n"
  }

  close(ch)
  return msg, nil
}
