package main

import (
	"github.com/jamm3e3333/quiz-app/cli"
	"github.com/jamm3e3333/quiz-app/config"
	"github.com/jamm3e3333/quiz-app/grpc"
)

func main() {
	cfg, err := config.NewParseConfigForENV()
	if err != nil {
		panic(err)
	}

	client := grpc.NewClient(uint32(cfg.GRPCServer.Port))

	submitCMD := cli.NewSubmitCMD(client)
	quizCMD := cli.NewQuizCMD(client, submitCMD)

	rootCLI := cli.InitCLI()
	rootCLI.RegisterCMDS(quizCMD, submitCMD)
	rootCLI.Run()
}
