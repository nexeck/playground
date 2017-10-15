package cmd

import (
	"fmt"
	"github.com/nexeck/playground/pkg/version"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(cmdVersion())
}

// NewCmdVersion adds version command
func cmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
}

func printVersion() {
	fmt.Println("---------------------------------------------------")
	fmt.Printf("AppVersion: %s\n", version.Info.AppVersion)
	fmt.Printf("Build: %s\n", version.Info.BuildDate)
	fmt.Printf("GitHash: %s\n", version.Info.GitCommit)
	fmt.Printf("GitBranch: %s\n", version.Info.GitBranch)
	fmt.Printf("GitState: %s\n", version.Info.GitState)
	fmt.Println("---------------------------------------------------")
}
