package cmdline

import (
  "strings"
  "fmt"
  "io"
  "bufio"

  "playground/master/client"
  . "playground/master/ferrors"
)

func ShowPrompt(clState *client.State) {
  addr, many := "", ""
  if clState.Connected {
    addr = clState.GetMainAddr().String()
    if clState.NConnections > 1 {
      many = "*"
    }
  }
  prompt := fmt.Sprintf("[%s%s]> ", addr, many)
  fmt.Print(prompt)
}

func ReadLine(reader *bufio.Reader) (string, error) {
  line, err := reader.ReadString('\n')
  if err == io.EOF {
    return "exit", nil
  }

  if err != nil {
    return "", MakeError(ErrCmdline, err.Error())
  }
  return line, nil
}

func ParseLine(line string) (Command, []string) {
  words := strings.Fields(line)

  for i, str := range commandString {
    if words[0] == str {
      return Command(i), words[1:]
    }
  }
  return Unknown, words[1:]
}

