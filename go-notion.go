package main

import (
	"fmt"
	"go-notion/pkg/notionjson"
	"os"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the database",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken, databaseId := notionjson.GetAccessTokens()
		response := notionjson.QueryDatabase(databaseId, apiToken)
		fmt.Fprintln(os.Stderr, "=== QueryDatabase response:")
		fmt.Println(string(response))
	},
}

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve data from the database",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken, databaseId := notionjson.GetAccessTokens()
		response := notionjson.RetrieveDatabase(databaseId, apiToken)
		fmt.Fprintln(os.Stderr, "=== RetrieveDatabase response:")
		fmt.Println(string(response))
	},
}

func main() {
	rootCmd := &cobra.Command{Use: "go-notion"}
	rootCmd.AddCommand(queryCmd, retrieveCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Execute()
}

