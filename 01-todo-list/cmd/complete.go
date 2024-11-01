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

var done = map[bool]string{true: "DONE", false: "UNDONE"}

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "This command mark the task as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parseInt, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		task, err := markAsDoneTask(parseInt)
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
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
