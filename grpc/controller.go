package grpc

import (
	"context"

	"github.com/jamm3e3333/quiz-app/app/quiz"
	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"
	"github.com/jamm3e3333/quiz-app/model"
)

type QuizHandler interface {
	ListQuestions(ctx context.Context) ([]model.Question, error)
	SubmitQuiz(ctx context.Context, questionIDToAnswerIndex map[int32]int32) (quiz.SubmissionDTO, error)
}

type ResultCalculator interface {
}

type QuizController struct {
	proto.UnimplementedQuizServiceServer

	quizHan QuizHandler
}

func NewQuizController(quizHan QuizHandler) *QuizController {
	return &QuizController{quizHan: quizHan}
}

func (c *QuizController) ListQuestions(ctx context.Context, _ *proto.ListQuestionsRequest) (*proto.ListQuestionsResponse, error) {
	q, err := c.quizHan.ListQuestions(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*proto.ListQuestion, len(q))
	for i, question := range q {
		res[i] = &proto.ListQuestion{
			Id:       question.ID,
			Question: question.Text,
			Answers:  question.Answers,
		}
	}

	return &proto.ListQuestionsResponse{Questions: res}, nil
}

func (c *QuizController) SubmitQuiz(ctx context.Context, req *proto.SubmitQuizRequest) (*proto.SubmitQuizResponse, error) {
	questionIDToAnswerIndex := make(map[int32]int32)
	for _, a := range req.GetAnswers() {
		questionIDToAnswerIndex[a.GetQuestionId()] = a.GetAnswerIndex()
	}

	sub, err := c.quizHan.SubmitQuiz(ctx, questionIDToAnswerIndex)
	if err != nil {
		return nil, err
	}

	return &proto.SubmitQuizResponse{
		SuccessRatePercentage: sub.SuccessRatePercentage,
		CorrectAnswers:        sub.CorrectA,
		TotalQuestions:        sub.TotalQ,
	}, nil
}
