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
		response, error := internal.CreateMessage(username, content)

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

var likeCmd = &cobra.Command{
	Use:   "like [messageID]",
	Short: "Like a message by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		messageID := args[0]
		response, error := internal.LikeMessage(messageID)

		if error != nil {
			fmt.Printf("Cannot register user.\n")
			panic(error)
		}

		if response.StatusCode == 200 {
			fmt.Printf("Message '%s' liked successfully.\n", messageID)
		} else {
			fmt.Printf("Failed to like message '%s'.\n", messageID)
		}
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(likeCmd)
}
