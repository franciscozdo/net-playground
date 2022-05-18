package ferrors

import (
  "fmt"
)

type Module int

const (
  ErrMain Module = iota
  ErrClient
  ErrCmdline
  ErrConfig
  ErrTests
)

var moduleName = [...]string{
  "main",
  "client",
  "cmdline",
  "config",
  "tests"}

type fError struct {
  Where Module
  Msg string
}

func (m Module) String() string {
  return moduleName[int(m)]
}

func (e *fError) Error() string {
  if e == nil {
    return ""
  }
  return fmt.Sprintf("[%s] %s", e.Where, e.Msg)
}

func MakeError(m Module, msg string) *fError {
  return &fError{m, msg}
}
