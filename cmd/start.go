package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
	"socialite/controllers"
	"socialite/ent"
	"socialite/services"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts Socialite",
	Long:  "Starts Socialite server",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading environment variables:", err.Error())
		}

		err, meili := services.CreateMeiliClient(os.Getenv("MEILISEARCH_URL"), "")
		if err != nil {
			log.Fatal("Failed creating a meili client:", err.Error())
		}
		ent.InitProductionDatabase(os.Getenv("DATABASE_URL"))
		controllers.StartServer(os.Getenv("PORT"), ent.Database(), meili)
	},
}
