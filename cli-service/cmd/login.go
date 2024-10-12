package cmd

import (
	"cli-service/internal"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login [username]",
	Short: "Log in as a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		apiClient := internal.NewAPIClient()
		response, err := apiClient.LoginUser(username)
		if err != nil {
			fmt.Printf("Error logging in: %v\n", err)
			return
		}

		if response.StatusCode == http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Printf("User '%s' logged in: %s\n", username, string(body))
		} else {
			fmt.Printf("Failed to log in user '%s'. Status code: %d\n", username, response.StatusCode)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
