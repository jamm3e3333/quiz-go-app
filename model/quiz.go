package model

type Question struct {
	ID      int32
	Text    string
	Answers []string
	Correct int32
}

type QuizModel struct {
	questions []Question
}

func NewQuizModel() *QuizModel {
	return &QuizModel{
		questions: []Question{
			{
				ID:      0,
				Text:    "What is the capital of France?",
				Answers: []string{"Berlin", "Madrid", "Paris", "Rome"},
				Correct: 2,
			},
			{
				ID:      1,
				Text:    "Which planet is known as the Red Planet?",
				Answers: []string{"Earth", "Mars", "Jupiter", "Saturn"},
				Correct: 1,
			},
		},
	}
}

func (q *QuizModel) Questions() []Question {
	return q.questions
}
