package cmdline

type Command int64

const (
  Unknown Command = iota
  Exit
  ShowHelp
  SchedTest
  ShowConfig
  Connect
  ShowConn
  SwitchConn
  Disconnect
  TestAll
)

var commandString = [...]string{
  "unknown command",
  "exit",
  "help",
  "test",
  "config",
  "connect",
  "connections",
  "switch",
  "disconnect",
  "testall",
  "" }

var helpMessage string =
`Available commands:
  config                   -- show known hosts you can connect to
  connect [host] [port]    -- connect to a host on a given port
                              (e.g. connect 123.4.5.6 50009)
  disconnect               -- disconnect with current host
  connections              -- show connected hosts
  switch [id]              -- change current host
  test [type]              -- schedule test on host connected to
                              (e.g. schedtest echo)
  testall [type]           -- run test on all connected host at the time
  exit                     -- leave the program
  help                     -- show this message
`
