package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const banner = "  ____             _                               \n" +
	" |  _ \\  ___   ___| | _____ _ __ __ _  ___ _ __   \n" +
	" | | | |/ _ \\ / __| |/ / _ \\ '__/ _` |/ _ \\ '_ \\  \n" +
	" | |_| | (_) | (__|   <  __/ | | (_| |  __/ | | | \n" +
	" |____/ \\___/ \\___|_|\\_\\___|_|  \\__, |\\___|_| |_| \n" +
	"                                |___/              \n"

var bannerShown bool

var rootCmd = &cobra.Command{
	Use:   "dockergen",
	Short: "Generate Dockerfile/Containerfile templates",
	Long:  "dockergen is a CLI for generating optimized Dockerfile/Containerfile templates for common runtimes.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if bannerShown {
			return
		}
		fmt.Fprint(os.Stdout, banner)
		bannerShown = true
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpTemplate(banner + "\n" + rootCmd.HelpTemplate())
	rootCmd.AddCommand(autogenCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}
