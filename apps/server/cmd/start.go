package cmd

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os"
	"socialite/controllers"
	"socialite/models"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts Socialite",
	Long:  "Starts Socialite client and server",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading environment variables:", err.Error())
		}

		controllers.StartServer(os.Getenv("PORT"))
		models.MigrateDatabase(models.InitDatabase(os.Getenv("DATABASE_URL")))
	},
}
