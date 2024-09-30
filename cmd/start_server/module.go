package main

import (
	"github.com/jamm3e3333/quiz-app/app/quiz"
	"github.com/jamm3e3333/quiz-app/app/submission"
	"github.com/jamm3e3333/quiz-app/grpc"
	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"
	"github.com/jamm3e3333/quiz-app/model"
)

func RegisterModule(g *grpc.Server) {
	quizHan := quiz.NewHandler(
		model.NewQuizModel(),
		model.NewSubmissionModel(),
		submission.NewResultCalculator(),
	)
	quizCTRL := grpc.NewQuizController(quizHan)

	proto.RegisterQuizServiceServer(g.Gs, quizCTRL)
}
