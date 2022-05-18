package server

import (
  "log"
  "fmt"
  "net"
  "sync"

  "google.golang.org/grpc"

  "playground/host/config"
  "playground/host/tests"
  msg "playground/host/messages"
)

type ServerRoutineChannels struct {
  DataReq, TestReq, NewTest chan<- tests.TestEntry
  DataRes <-chan tests.TestEntry
}

func ServerRoutine(chs ServerRoutineChannels, wg sync.WaitGroup, conf *config.Config) {
  wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatal(err)
	}

  serv := &Server{channels: chs}
	grpcServer := grpc.NewServer()

	msg.RegisterHostServer(grpcServer, serv)

  log.Printf("Starting server at port %d...", conf.Port)

  err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
