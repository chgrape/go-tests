/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"main/db"
	"time"

	"github.com/spf13/cobra"
)

var title string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Long:  `Use the flag -t to add a description to the task`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := db.Con.Exec("INSERT INTO notes (title, created_at) VALUES (?, ?)",
			title,
			time.Now(),
		)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "The title of the task")
	addCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
