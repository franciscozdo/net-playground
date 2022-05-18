package client

import (
  "fmt"
  "net"
  "context"
  "time"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"

  msg "playground/master/messages"
  . "playground/master/ferrors"
)

type Client struct {
  Conn *grpc.ClientConn
  Cl msg.HostClient
  Addr net.IP
}

func makeConnection(host string, port int) (*grpc.ClientConn, error) {
  ctx, _ := context.WithTimeout(context.Background(), 3 * time.Second)
	return grpc.DialContext(ctx, fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func makeClient(host net.IP, port int) (*Client, error) {
  conn, err := makeConnection(host.String(), port)
  if err != nil {
    return nil, err
  }
  cl := Client{Conn: conn, Cl: msg.NewHostClient(conn)}
  cl.Addr = host

  /* Send hello message to host */
  _, err = cl.Cl.Hello(context.Background(), &msg.HelloRequest{Name: "master"})
  if err != nil {
    return nil, err
  }
  return &cl, nil
}

func (c *Client) RunTest(id int32, ttype msg.Test) (error) {
  testR := msg.TestRequest{Id: id, TestType: ttype}
  resp, err := c.Cl.RunTest(context.Background(), &testR)
  if err != nil {
    return MakeError(ErrClient, err.Error())
  }
  if resp.GetScheduled() == false {
    return MakeError(ErrClient, "test not scheduled")
  }
  return nil
}

func (c *Client) GetResult(id int32) (bool, []string, error) {
  resR := msg.ResultRequest{Id: id}
  resp, err := c.Cl.GetResult(context.Background(), &resR)
  if err != nil {
    return false, nil, MakeError(ErrClient, err.Error())
  }
  t := resp.GetReachable()
  done := resp.GetSuccess()
  return done, t, nil
}

func (c *Client) Close() error {
  return c.Conn.Close()
}
