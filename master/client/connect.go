package client

import (
  "net"
  "fmt"

  . "playground/master/ferrors"
)

func (s *State) Connect(host net.IP, port int) error {

  if s.NConnections == s.MaxConnections {
    msg := fmt.Sprintf("tried to make to many connections (only %d are allowed)", s.MaxConnections)
    return MakeError(ErrClient, msg)
  }

  if s.IsConnected(host) {
    return MakeError(ErrClient, "already connected")
  }
  cl, err := makeClient(host, port)
  if err != nil {
    return MakeError(ErrClient, err.Error())
  }

  s.AddClient(cl)

  return nil
}

func (s *State) Disconnect() error {
  cl := s.GetMainConnection()
  err := cl.Close()
  if err != nil {
    return err
  }
  s.RemoveClient(s.MainConnection)
  return nil
}
