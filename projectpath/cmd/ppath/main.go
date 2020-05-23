package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"aduu.dev/utils/completion"
	"aduu.dev/utils/projectpath"
)

// RootCMD creates a root command of a program.
func RootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ppath",
		Short: "prints the current default project path.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(projectpath.ProjectPath())
			return nil
		},
	}
	cmd.SetOutput(os.Stdout) // Else printing to stderr with cmd.Print().
	cmd.AddCommand(completion.NewCompletionCMD())
	cmd.AddCommand()
	return cmd
}

func main() {
	if err := RootCMD().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
