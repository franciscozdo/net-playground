package cmdline

import (
  "strconv"

  msg "playground/master/messages"
  . "playground/master/ferrors"
)

func ConnectArgs(args []string) (string, error) {
  if len(args) < 1 {
    return "", MakeError(ErrCmdline, "connect - to few arguments")
  }

  host := args[0]

  return host, nil
}

func SchedTestArgs(args []string) (msg.Test, error) {
  if len(args) < 1 {
    return 0, MakeError(ErrCmdline, "sched test - to few arguments")
  }

  var testType msg.Test
  switch args[0] {
  case "echo":
    testType = msg.Test_ECHO
  case "http":
    testType = msg.Test_HTTP
  default:
    testType = msg.Test_UNKNOWN
  }
  return testType, nil
}

func SwitchConnArgs(args []string) (int, error) {
  if len(args) < 1 {
    return 0, MakeError(ErrCmdline, "switch - you have to specify id of new current connection")
  }

  id, e := strconv.Atoi(args[0])
  if e != nil {
    return 0, MakeError(ErrCmdline, "swith - argument must be a number")
  }
  return id, nil
}

func TestAllArgs(args []string) (msg.Test, error) {
  if len(args) < 1 {
    return 0, MakeError(ErrCmdline, "test all - to few arguments")
  }

  var testType msg.Test
  switch args[0] {
  case "echo":
    testType = msg.Test_ECHO
  case "http":
    testType = msg.Test_HTTP
  default:
    testType = msg.Test_UNKNOWN
  }
  return testType, nil
}
