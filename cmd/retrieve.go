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

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve the database metadata",
	Long:  `Retrieve the metadata for the Notion database specified in the .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiToken, databaseId := notionjson.GetAccessTokens()
		response := notionjson.RetrieveDatabase(databaseId, apiToken)
		fmt.Fprintln(os.Stderr, "=== RetrieveDatabase response:")
		fmt.Println(string(response))
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbdataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbdataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
