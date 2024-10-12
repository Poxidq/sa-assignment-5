package cmd

import (
	"cli-service/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "Get the latest 10 messages",
	Run: func(cmd *cobra.Command, args []string) {
		client := internal.NewAPIClient()
		feed, err := client.GetFeed()

		if err != nil {
			fmt.Printf("Error retrieving the feed: %v\n", err)
			return
		}

		if len(feed) > 0 {
			fmt.Println("Latest 10 messages:")
			for _, message := range feed {
				fmt.Printf("ID: %d | User: %s | Likes: %d\nMessage: %s\n\n", message.ID, message.Username, message.Likes, message.Content)
			}
		} else {
			fmt.Println("No messages found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(feedCmd)
}
