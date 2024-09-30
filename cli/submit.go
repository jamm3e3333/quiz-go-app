package cli

import (
	"encoding/json"
	"fmt"

	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"

	"github.com/spf13/cobra"
)

type SubmitHandler interface {
	SubmitQuiz(
		request *proto.SubmitQuizRequest,
		md map[string]string,
	) (*proto.SubmitQuizResponse, error)
}

func NewSubmitCMD(client SubmitHandler) *cobra.Command {
	return &cobra.Command{
		Use:   "submit",
		Short: "Submit quiz answers",
		Run: func(cmd *cobra.Command, args []string) {
			var answers map[int32]int32
			err := json.Unmarshal([]byte(args[0]), &answers)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Failed to unmarshal user answers: %v\n", err)
				return
			}

			submittedAnswers := make([]*proto.SubmittedAnswer, len(answers))

			for k, v := range answers {
				submittedAnswers[k] = &proto.SubmittedAnswer{
					QuestionId:  k,
					AnswerIndex: v,
				}
			}

			res, err := client.SubmitQuiz(&proto.SubmitQuizRequest{Answers: submittedAnswers}, nil)
			if err != nil {
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Failed to submit quiz: %v\n", err)
				return
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "You answered %d out of %d questions correctly.\n", res.GetCorrectAnswers(), res.GetTotalQuestions())

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "You were better than %.2f%% of all quizzers.\n", res.GetSuccessRatePercentage())
		},
	}

}
