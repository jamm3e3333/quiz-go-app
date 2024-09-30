package submission

import "github.com/jamm3e3333/quiz-app/model"

type ResultCalculator struct{}

func NewResultCalculator() *ResultCalculator {
	return &ResultCalculator{}
}

func (s *ResultCalculator) CalculateResults(questions []model.Question, questionIDToAnswerID map[int32]int32) (correctAnswers, totalQuestions int32, percentage float32) {
	totalQuestions = int32(len(questions))
	correctAnswers = 0

	for i, question := range questions {
		if answer, ok := questionIDToAnswerID[int32(i)]; ok && answer == question.Correct {
			correctAnswers++
		}
	}

	percentage = (float32(correctAnswers) / float32(totalQuestions)) * 100

	return correctAnswers, totalQuestions, percentage
}

func (s *ResultCalculator) CompareWithPreviousSubmissions(submissions []model.Submission, currentScorePercentage float32) (successRatePercentage float32) {
	if len(submissions) == 0 {
		return 100.0
	}

	betterThanCount := 0
	for _, submission := range submissions {
		previousPercentage := (float32(submission.CorrectAnswers) / float32(submission.TotalQuestions)) * 100
		if currentScorePercentage >= previousPercentage {
			betterThanCount++
		}
	}

	successRatePercentage = (float32(betterThanCount) / float32(len(submissions))) * 100

	return successRatePercentage
}
