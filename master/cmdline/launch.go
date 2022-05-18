package cmdline

import (
  "fmt"
  "os"

  "playground/master/client"
  "playground/master/config"
  "playground/master/tests"
  msg "playground/master/messages"
)

func EvalCommand(cmd Command, args []string, cl *client.State, conf *config.Config) error {
  switch cmd {
  case Unknown:
    fmt.Println("Unknown command. Type h to see all commands.")
  case ShowHelp:
    fmt.Print(helpMessage)
  case ShowConfig:
    conf.Print()
  case Connect:
    hostname, err := ConnectArgs(args)
    if err != nil {
      return err
    }

    host, port, err := conf.GetAddress(hostname)
    if err != nil {
      return err
    }

    err = cl.Connect(host, port)
    if err != nil {
      return err
    }
  case Disconnect:
    cl.Disconnect()
  case ShowConn:
    msg := cl.ShowConnections(conf)
    fmt.Println("Active connections:")
    fmt.Print(msg)
  case SwitchConn:
    id, err := SwitchConnArgs(args)
    if err != nil {
      return err
    }
    err = cl.Switch(id)
    return err
  case SchedTest:
    test, err := SchedTestArgs(args)
    if err != nil {
      return err
    }
    err = ScheduleTest(test, cl, conf)
    return err
  case TestAll:
    test, err := SchedTestArgs(args)
    if err != nil {
      return err
    }
    err = SchedulaManyTests(test, cl, conf)
    return err
  case Exit:
    os.Exit(0)
  }
  return nil
}

func ScheduleTest(ttype msg.Test, clState *client.State, conf *config.Config) error {
  cl := clState.GetMainConnection()

  res, err := tests.TestOne(ttype, cl, conf)
  if err != nil {
    return err
  }

  fmt.Printf("Result of %v test:\n", ttype)
  fmt.Println(res)

  return nil
}

func SchedulaManyTests(ttype msg.Test, clState *client.State, conf *config.Config) error {
  res, err := tests.TestMany(ttype, clState, conf)

  if err != nil {
    return err
  }

  fmt.Printf("Results of %v test for all hosts:\n", ttype)
  fmt.Println(res)

  return nil
}

