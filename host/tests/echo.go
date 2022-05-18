package tests

import (
  "net"
  "bufio"
  "strconv"
  "errors"
  "time"
  "log"
)

func connect(host net.IP, port int) (net.Conn, error) {
  addr := host.String()
  ports := strconv.Itoa(port)
  d := net.Dialer{Timeout: 3 * time.Second}
  conn, err := d.Dial("tcp", net.JoinHostPort(addr, ports))

  if err != nil {
    return nil, err
  }
  return conn, nil
}

var testString string = "Echo test string.\n"

func EchoTest(host net.IP) error {
  conn, err := connect(host, 7)
  if err != nil { return err }

  reader := bufio.NewReader(conn)

  _, err = conn.Write([]byte(testString))
  if err != nil { return err }

  res, err := reader.ReadString('\n')
  if err != nil { return err }

  log.Print("'%s' ?= '%s'", res, testString)
  if res != testString {
    return errors.New("echo test failed")
  }

  return nil
}
