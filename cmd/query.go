/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-notion/pkg/notionjson"
	"os"

	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the database to get the list of data for each page",
	Long: `Query the Notion database specified in the .env file
to get the list of data for each page.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken, databaseId := notionjson.GetAccessTokens()
		response := notionjson.QueryDatabase(databaseId, apiToken)
		fmt.Fprintln(os.Stderr, "=== QueryDatabase response:")
		fmt.Println(string(response))
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pagelistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pagelistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
