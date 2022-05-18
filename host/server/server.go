package server

import (
	"log"

	"golang.org/x/net/context"
  "google.golang.org/grpc/peer"

  msg "playground/host/messages"
  "playground/host/tests"
)

type Server struct {
  /* channels to communicate with other routines */
  channels ServerRoutineChannels

  /* needed by grpc */
  msg.UnimplementedHostServer
}

func (s *Server) Hello(ctx context.Context, in *msg.HelloRequest) (*msg.HelloResponse, error) {
  p, _ := peer.FromContext(ctx)
  log.Printf("Hello from %s (%s)", in.Name, p.Addr.String())
  return &msg.HelloResponse{Name: "host"}, nil
}

func (s *Server) RunTest(ctx context.Context, in *msg.TestRequest) (*msg.TestRequestResp, error) {
  log.Printf("Request to start test %s", in.TestType.String())
  entry := tests.TestEntry{Id: int(in.Id), Type: in.TestType, Done: false}
  s.channels.NewTest <- entry
  s.channels.TestReq <- entry
  log.Printf("Recorded test %d", in.Id)
  return &msg.TestRequestResp{Id: in.Id, Scheduled: true}, nil
}

func (s *Server) GetResult(ctx context.Context, in *msg.ResultRequest) (*msg.ResultData, error) {
  id := int(in.Id)
  log.Printf("Request for data from test %d", id)

  s.channels.DataReq <- tests.TestEntry{Id: id}
  t := <-s.channels.DataRes

  log.Printf("Test done: %v", t.Done)
  var r []string
  if t.Done {
    r = *t.Data
  }

  m := msg.ResultData{Id: int32(id), Reachable: r, Success: t.Done}
  return &m, nil
}

/*
  rpc RunTest(TestRequest) returns (TestRequestResp) {}
  rpc GetResult(ResultRequest) returns (ResultData) {}
*/


