/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/db"

	"github.com/spf13/cobra"
)

var id int

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove task",
	Long:  `Pass id with -i flag to remove task. If id doesn't exist, the program exits`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := db.Con.Exec("DELETE FROM notes WHERE id=?", id)

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

		fmt.Println("Deleted todo with id", id)
	},
}

func init() {
	rmCmd.Flags().IntVarP(&id, "id", "i", 0, "Id of task to be removed")
	rmCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
