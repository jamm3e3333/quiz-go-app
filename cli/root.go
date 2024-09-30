package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type CLI struct {
	rootCMD *cobra.Command
}

func InitCLI() *CLI {
	rootCmd := &cobra.Command{
		Use:   "quiz",
		Short: "CLI Quiz Application",
		Long:  `A CLI application that allows users to take quizzes and see how they compare to others.`,
	}

	return &CLI{
		rootCMD: rootCmd,
	}
}

func (c *CLI) RegisterCMDS(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		c.rootCMD.AddCommand(cmd)
	}
}

func (c *CLI) Run() {
	if err := c.rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
