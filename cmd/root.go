package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "echoer",
	Version: "0.0.1",
	Short:   "Echoer",
	Long: `Echoer
	
	Echoes back to the caller the string it receives.`,

	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("server failure: %v", err)
	}
}
