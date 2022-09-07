package cmd

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"socialite/models"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "This command allows you to create fake database models filled with data. Use for development purposes only.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must specify database model")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		model := args[0]
		switch model {
		case "user":
			bytes, err := json.Marshal(models.GenerateUser())
			if err != nil {
				cmd.PrintErr("Error:", err)
				return
			}
			cmd.Println(string(bytes))
		}
	},
}
