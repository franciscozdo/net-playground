package client

import (
  "net"
  "fmt"

  . "playground/master/ferrors"
  "playground/master/config"
)

type State struct {
  Connected bool
  NConnections int
  MaxConnections int
  MainConnection int // idx in array of connections
  Connections []*Client
}

func Init(maxClients int) (*State, error) {
  return &State{false, 0, maxClients, 0, make([]*Client, maxClients)}, nil
}

func (s *State) GetMainAddr() net.IP {
  c := s.GetMainConnection()
  if c == nil {
    return nil
  }
  return c.Addr
}

func (s *State) GetMainConnection() *Client {
  if !s.Connected {
    return nil
  }
  return s.Connections[s.MainConnection]
}

func (s *State) IsConnected(addr net.IP) bool {

  for i := 0; i < s.NConnections; i++ {
    cl := s.Connections[i]
    if cl.Addr.Equal(addr) {
      return true
    }
  }
  return false
}

func (s *State) ShowConnections(conf *config.Config) string {
  msg := ""
  for i := 0; i < s.NConnections; i++ {
    cl := s.Connections[i]
    a := cl.Addr
    nm, _ := conf.GetName(a)
    msg += fmt.Sprintf("  %d: %s(%s)\n", i, nm, a)
  }
  return msg
}

func (s *State) Switch(id int) error {
  if id >= s.NConnections {
    return MakeError(ErrClient, "index of connection is too large")
  }
  s.MainConnection = id
  return nil
}

func (s *State) AddClient(cl *Client) {
  idx := s.NConnections
  s.Connections[idx] = cl
  s.MainConnection = idx
  s.NConnections++
  s.Connected = true
}

func (s *State) RemoveClient(idx int) {
  if s.NConnections == 0 {
    return
  }
  /* swap with last */
  last := s.NConnections - 1
  s.Connections[idx], s.Connections[last] = s.Connections[last], s.Connections[idx]

  /* remove last */
  s.Connections[last] = nil
  s.NConnections--

  /* assure main connection always exists */
  if s.MainConnection >= s.NConnections {
    s.MainConnection = 0
  }

  if s.NConnections == 0 {
    s.Connected = false
  }
}
