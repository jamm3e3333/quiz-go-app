package cli

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"
)

// MockQuizProvider simulates the QuizProvider behavior
type MockQuizProvider struct {
	Response *proto.ListQuestionsResponse
	Err      error
}

func (m *MockQuizProvider) ListQuestions(md map[string]string) (*proto.ListQuestionsResponse, error) {
	return m.Response, m.Err
}

func TestQuizCommand_Success(t *testing.T) {
	mockQuizProvider := &MockQuizProvider{
		Response: &proto.ListQuestionsResponse{
			Questions: []*proto.ListQuestion{
				{
					Question: "What is the capital of France?",
					Answers:  []string{"Berlin", "Madrid", "Paris", "Rome"},
				},
				{
					Question: "Which planet is known as the Red Planet?",
					Answers:  []string{"Earth", "Mars", "Jupiter", "Saturn"},
				},
			},
		},
		Err: nil,
	}

	mockSubmitHandler := &MockSubmitHandler{
		Response: &proto.SubmitQuizResponse{
			CorrectAnswers:        2,
			TotalQuestions:        2,
			SuccessRatePercentage: 80.0,
		},
		Err: nil,
	}

	submitCmd := NewSubmitCMD(mockSubmitHandler)

	cmd := NewQuizCMD(mockQuizProvider, submitCmd)

	output := &bytes.Buffer{}
	cmd.SetOut(output)

	userInput := strings.NewReader("3\n2\n")
	cmd.SetIn(userInput)

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `Question 1: What is the capital of France?
  1. Berlin
  2. Madrid
  3. Paris
  4. Rome
Your answer: Question 2: Which planet is known as the Red Planet?
  1. Earth
  2. Mars
  3. Jupiter
  4. Saturn
Your answer: You answered 2 out of 2 questions correctly.
You were better than 80.00% of all quizzers.
`

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestQuizCommand_FetchQuestionsFailure(t *testing.T) {
	mockQuizProvider := &MockQuizProvider{
		Response: nil,
		Err:      errors.New("failed to fetch questions"),
	}

	mockSubmitHandler := &MockSubmitHandler{}

	submitCmd := NewSubmitCMD(mockSubmitHandler)
	cmd := NewQuizCMD(mockQuizProvider, submitCmd)

	output := &strings.Builder{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Failed to fetch questions: failed to fetch questions\n"

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
