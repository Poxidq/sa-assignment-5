package cmd

import (
	"cli-service/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post [username] [message]",
	Short: "Post a new message",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		content := args[1]

		client := internal.NewAPIClient()
		response, error := client.CreateMessage(username, content)

		if error != nil {
			fmt.Printf("Cannot register user.\n")
			panic(error)
		}

		if response != nil && response.StatusCode == 201 {
			fmt.Println("Message posted successfully.")
		} else {
			fmt.Println("Failed to post message.")
		}
	},
}
