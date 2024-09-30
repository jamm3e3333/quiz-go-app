package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"

	proto "github.com/jamm3e3333/quiz-app/grpc/protobuff"
	"github.com/spf13/cobra"
)

type QuizProvider interface {
	ListQuestions(
		md map[string]string,
	) (*proto.ListQuestionsResponse, error)
}

func NewQuizCMD(client QuizProvider, submitCMD *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:   "quiz",
		Short: "Start a new quiz",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := client.ListQuestions(nil)
			if err != nil {
				_, _ = fmt.Fprintln(cmd.OutOrStdout(), "Failed to fetch questions:", err)
				return
			}

			userAnswers := make(map[int32]int32)

			scanner := bufio.NewScanner(cmd.InOrStdin())

			for i, question := range res.GetQuestions() {
				_, err = fmt.Fprintf(cmd.OutOrStdout(), "Question %d: %s\n", i+1, question.Question)
				if err != nil {
					return
				}

				for j, answer := range question.Answers {
					_, err = fmt.Fprintf(cmd.OutOrStdout(), "  %d. %s\n", j+1, answer)
					if err != nil {
						return
					}
				}
				_, err = fmt.Fprint(cmd.OutOrStdout(), "Your answer: ")
				if err != nil {
					return
				}

				if !scanner.Scan() {
					_, err = fmt.Fprintln(cmd.OutOrStderr(), "Failed to read input")
					if err != nil {
						return
					}
					return
				}

				answerStr := scanner.Text()
				answerIdx, err := strconv.Atoi(answerStr)

				if err != nil || answerIdx < 1 || answerIdx > len(question.Answers) {
					_, err = fmt.Fprintln(cmd.OutOrStdout(), "Invalid answer. Please choose a valid number.")
					if err != nil {
						return
					}
					i--
					continue
				}
				userAnswers[int32(i)] = int32(answerIdx) - 1
			}

			data, err := json.Marshal(&userAnswers)
			if err != nil {
				_, err = fmt.Fprintln(cmd.OutOrStdout(), "Failed to marshal user answers:", err)
				if err != nil {
					return
				}
				return
			}

			submitCMD.Run(cmd, []string{string(data)})
		},
	}
}
