/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var undoneCmd = &cobra.Command{
	Use:   "undone",
	Short: "This command mark a complete task as undone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parseInt, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		task, err := changeStatusTask(parseInt, false)
		if err != nil {
			log.Fatal(err)
		}
		newTable := table.New(os.Stdout)
		newTable.SetHeaders("ID", "Name", "Created", "Status")
		newTable.SetHeaderStyle(table.StyleBold)
		newTable.SetLineStyle(table.StyleBlue)
		newTable.SetDividers(table.UnicodeRoundedDividers)
		newTable.AddRow(strconv.Itoa(task.id), task.name, timediff.TimeDiff(task.created), done[task.isComplete])
		newTable.Render()
	},
}

func init() {
	rootCmd.AddCommand(undoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
