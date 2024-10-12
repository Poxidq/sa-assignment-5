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
		feed, error := internal.GetFeed()

		if error != nil {
			fmt.Printf("Cannot register user.\n")
			panic(error)
		}

		if len(feed) > 0 {
			fmt.Println("Latest 10 messages:")
			for _, message := range feed {
				fmt.Printf("User: %s\nMessage: %s\n\n", message.Username, message.Content)
			}
		} else {
			fmt.Println("No messages found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(feedCmd)
}
