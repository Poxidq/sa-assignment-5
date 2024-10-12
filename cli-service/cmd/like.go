package cmd

import (
	"fmt"
	"net/http"
	"twitter-cli/internal"

	"github.com/spf13/cobra"
)

var likeCmd = &cobra.Command{
	Use:   "like [messageID]",
	Short: "Like a message by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		messageID := args[0]
		apiClient := internal.NewAPIClient()
		response, err := apiClient.LikeMessage(messageID)
		if err != nil {
			fmt.Printf("Error liking message: %v\n", err)
			return
		}

		if response.StatusCode == http.StatusOK {
			fmt.Printf("Message '%s' liked successfully.\n", messageID)
		} else {
			fmt.Printf("Failed to like message '%s'. Status code: %d\n", messageID, response.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(likeCmd)
}
