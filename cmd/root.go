package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "socialite",
	Short: "Socialite is an open-source social media platform",
	Long: "Socialite is an open-source social media platform built with Nuxt.js, Go, CockroachDB, and Meilisearch. " +
		"You can use it to spin up your own social media in minutes. This command is an entrypoint to the Socialite CLI.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
