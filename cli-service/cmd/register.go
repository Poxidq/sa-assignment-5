package cmd

import (
	"cli-service/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "Register a new user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		response, error := internal.RegisterUser(username)
		if error != nil {
			fmt.Printf("Cannot register user.\n")
			panic(error)
		}
		if response.StatusCode == 201 {
			fmt.Printf("User '%s' registered successfully.\n", username)
		} else {
			fmt.Printf("Failed to register user '%s'.\n", username)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
