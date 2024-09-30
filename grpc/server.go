package grpc

import (
	"fmt"
	"net"
	"time"

	"github.com/jamm3e3333/quiz-app/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	KeepAliveTimePingInterval = 15 * time.Second
	KeepAlivePingTimeout      = 10 * time.Second
)

type Server struct {
	lg logger.Logger
	Gs *grpc.Server
	ec chan error
	p  uint32
}

func NewServer(logger logger.Logger, port uint32, useReflection bool, interceptors ...grpc.UnaryServerInterceptor) *Server {
	ka := keepalive.ServerParameters{
		Time:    KeepAliveTimePingInterval,
		Timeout: KeepAlivePingTimeout,
	}

	so := []grpc.ServerOption{
		grpc.KeepaliveParams(ka),
		grpc.ChainUnaryInterceptor(interceptors...),
	}

	gs := grpc.NewServer(so...)

	if useReflection {
		reflection.Register(gs)
	}

	return &Server{lg: logger, Gs: gs, ec: make(chan error), p: port}
}

func (s *Server) Run() <-chan error {
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.p))
		if err != nil {
			s.lg.Fatal(err)
		}
		defer func() {
			if s.ec != nil {
				close(s.ec)
			}
		}()

		s.ec <- s.Gs.Serve(l)
	}()

	return s.ec
}

func (s *Server) Shutdown() {
	s.Gs.GracefulStop()
}
