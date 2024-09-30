package cli

import (
	"bytes"
	"errors"
	"testing"

	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"
)

type MockSubmitHandler struct {
	Response *proto.SubmitQuizResponse
	Err      error
}

func (m *MockSubmitHandler) SubmitQuiz(
	_ *proto.SubmitQuizRequest,
	_ map[string]string,
) (*proto.SubmitQuizResponse, error) {
	return m.Response, m.Err
}

func TestSubmitCommand_Success(t *testing.T) {
	mockHandler := &MockSubmitHandler{
		Response: &proto.SubmitQuizResponse{
			CorrectAnswers:        2,
			TotalQuestions:        3,
			SuccessRatePercentage: 75.0,
		},
		Err: nil,
	}

	cmd := NewSubmitCMD(mockHandler)

	args := []string{`{"0":1,"1":2,"2":0}`}

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	cmd.SetArgs(args)

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "You answered 2 out of 3 questions correctly.\n" +
		"You were better than 75.00% of all quizzers.\n"

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestSubmitCommand_Failure(t *testing.T) {
	mockHandler := &MockSubmitHandler{
		Response: nil,
		Err:      errors.New("submission failed"),
	}

	cmd := NewSubmitCMD(mockHandler)

	args := []string{`{"0":1,"1":2,"2":0}`}

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	cmd.SetArgs(args)

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Failed to submit quiz: submission failed\n"

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}

func TestSubmitCommand_InvalidInput(t *testing.T) {
	mockHandler := &MockSubmitHandler{}

	cmd := NewSubmitCMD(mockHandler)

	args := []string{"invalid-json"}

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	cmd.SetArgs(args)

	err := cmd.Execute()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Failed to unmarshal user answers: invalid character 'i' looking for beginning of value\n"

	if output.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output.String())
	}
}
