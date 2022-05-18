package main

import (
  "fmt"
  "log"
  "bufio"
  "os"
  "flag"

  "playground/master/config"
  "playground/master/client"
  cmd "playground/master/cmdline"
)

var helloMessage string =
`Hello to master program.

You can test connection within all other hosts here.

Type h to see all commands.
`

func Exit() {
  /* TODO: clean after all actions (close connection, etc.) */

  fmt.Print("bye")
  os.Exit(0)
}

func ParseArgs() string {
  conf := flag.String("c", "config.yaml", "configuration file")
  flag.Parse()

  return *conf

}

func HandleErrorCrit(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func HandleError(err error) {
  if err != nil {
    log.Print(err)
  }
}

var clState *client.State
var conf *config.Config

func main() {
  configFile := ParseArgs()

  fmt.Print(helloMessage)

  /* INITIALIZE */
  var err error

  /* get config from file */
  conf, err = config.Init(configFile)
  HandleErrorCrit(err)

  conf.Print()

  /* init client */
  clState, err = client.Init(conf.MaxConnections)
  HandleErrorCrit(err)

  reader := bufio.NewReader(os.Stdin)

  /* MAIN LOOP */
  for {
    cmd.ShowPrompt(clState)
    line, err := cmd.ReadLine(reader)

    HandleError(err)

    /* normal cases */
    if line == "\n" {
      continue
    } else {
      command, args := cmd.ParseLine(line)

      if command == cmd.Exit {
        Exit()
      }

      err = cmd.EvalCommand(command, args, clState, conf)
      HandleError(err)
    }
  }
}
