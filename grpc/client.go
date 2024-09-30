package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	grpcTimeoutSec = 10 * time.Second
)

type Client struct {
	p uint32
}

func NewClient(p uint32) *Client {
	return &Client{
		p: p,
	}
}

func (c *Client) ListQuestions(
	md map[string]string,
) (*proto.ListQuestionsResponse, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", c.p), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	client := proto.NewQuizServiceClient(conn)

	meta := metadata.New(md)

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeoutSec)
	defer cancel()

	gCtx := metadata.NewOutgoingContext(ctx, meta)

	return client.ListQuestions(gCtx, &proto.ListQuestionsRequest{})
}

func (c *Client) SubmitQuiz(
	request *proto.SubmitQuizRequest,
	md map[string]string,
) (*proto.SubmitQuizResponse, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", c.p), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	client := proto.NewQuizServiceClient(conn)

	meta := metadata.New(md)

	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeoutSec)
	defer cancel()

	gCtx := metadata.NewOutgoingContext(ctx, meta)

	return client.SubmitQuiz(gCtx, request)
}
