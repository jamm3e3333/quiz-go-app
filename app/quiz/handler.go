package quiz

import (
	"context"
	"time"

	"github.com/jamm3e3333/quiz-app/model"
)

type QuestionsProvider interface {
	Questions() []model.Question
}

type SubmissionProvider interface {
	SaveSubmission(submission model.Submission)
	Submissions() []model.Submission
}

type SubmissionResultCalculator interface {
	CalculateResults(questions []model.Question, questionIDToAnswerID map[int32]int32) (correctAnswers, totalQuestions int32, percentage float32)
	CompareWithPreviousSubmissions(submissions []model.Submission, currentScorePercentage float32) (successRatePercentage float32)
}

type Handler struct {
	questionsProvider  QuestionsProvider
	submissionProvider SubmissionProvider
	submissionCalc     SubmissionResultCalculator
}

func NewHandler(questionsProvider QuestionsProvider, subProvider SubmissionProvider, subCalc SubmissionResultCalculator) *Handler {
	return &Handler{
		questionsProvider:  questionsProvider,
		submissionProvider: subProvider,
		submissionCalc:     subCalc,
	}
}

func (h *Handler) ListQuestions(ctx context.Context) ([]model.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	q := h.questionsProvider.Questions()

	return q, nil
}

type SubmissionDTO struct {
	SuccessRatePercentage float32
	CorrectA              int32
	TotalQ                int32
}

func (h *Handler) SubmitQuiz(ctx context.Context, questionIDToAnswerIndex map[int32]int32) (SubmissionDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	correctA, totalQ, percentage := h.submissionCalc.CalculateResults(h.questionsProvider.Questions(), questionIDToAnswerIndex)
	successRatePercentage := h.submissionCalc.CompareWithPreviousSubmissions(h.submissionProvider.Submissions(), percentage)

	h.submissionProvider.SaveSubmission(model.Submission{
		CorrectAnswers: correctA,
		TotalQuestions: totalQ,
	})

	return SubmissionDTO{
		SuccessRatePercentage: successRatePercentage,
		CorrectA:              correctA,
		TotalQ:                totalQ,
	}, nil
}
