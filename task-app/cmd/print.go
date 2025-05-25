/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/db"
	"main/models"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints the current tasks",
	Long:  `Outputs a table of the current tasks along with their ids, what they are, their status, when they were created and completed.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.Con.Query("SELECT * FROM notes")

		if err != nil {
			panic(err)
		}

		table := tablewriter.NewTable(os.Stdout)
		table.Header([]string{"id", "title", "status", "created_at", "completed_at"})

		for res.Next() {
			var l models.Todo
			err := res.Scan(&l.Id, &l.Title, &l.Completed, &l.Created_at, &l.Completed_at)
			if err != nil {
				panic(err)
			}

			comp_str := "-"
			status_str := "Not done"
			if l.Completed {
				status_str = "Done"
			}

			if l.Completed_at != nil {
				comp_str = l.Completed_at.Format("2006-01-02 15:04")
			}

			table.Append([]string{
				fmt.Sprintf("%d", l.Id),
				l.Title,
				status_str,
				l.Created_at.Format("2006-01-02 15:04"),
				comp_str,
			})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
