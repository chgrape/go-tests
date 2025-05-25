/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/db"
	"time"

	"github.com/spf13/cobra"
)

// tickCmd represents the tick command
var tickCmd = &cobra.Command{
	Use:   "tick",
	Short: "Mark as done",
	Long:  `Pass id with -i flag to mark task as done or revert to not being done. If id doesn't exist, the program exits`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.Con.Exec(`UPDATE notes 
					SET completed=NOT completed, completed_at=CASE 
						WHEN NOT completed 
						THEN ? 
						ELSE NULL 
					END WHERE id=?`, time.Now(), id)
		if err != nil {
			panic(err)
		}

		rows_affected, err := res.RowsAffected()

		if err != nil {
			panic(err)
		}

		if rows_affected == 0 {
			fmt.Println("Action Failed: No todo with id", id)
			return
		}
	},
}

func init() {
	tickCmd.Flags().IntVarP(&id, "id", "i", 0, "Id of task, which's status needs to be toggled")
	tickCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(tickCmd)
}
