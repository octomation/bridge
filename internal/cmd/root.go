package cmd

import "github.com/spf13/cobra"

// New returns the new root command.
func New() *cobra.Command {
	command := cobra.Command{
		Use:   "bridge",
		Short: "a bridge between the protobuf and the brief formats",
		Long:  "A bridge between the protobuf and the brief formats.",

		SilenceErrors: false,
		SilenceUsage:  true,
	}
	/* configure instance */
	command.AddCommand( /* related commands */ )
	return &command
}
