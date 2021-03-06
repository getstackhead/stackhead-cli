package module

import (
	"github.com/spf13/cobra"
)

func GetCommands() *cobra.Command {
	command := &cobra.Command{
		Use:   "module",
		Short: "StackHead module commands",
	}
	validate := Validate()
	command.AddCommand(validate)
	return command
}
