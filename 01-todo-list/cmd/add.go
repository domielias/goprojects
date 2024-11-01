/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var (
	Verbose bool
	Todo    string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command add a new task",
	Long:  "This command allow the user to add a new task to the list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rawTasks, err := getFileData()
		if err != nil {
			log.Fatal(err)
		}
		tasks, err := getListTasks(rawTasks)
		if err != nil {
			log.Fatal(err)
		}
		numbersRows := len(tasks)

		task := Task{numbersRows, args[0], time.Now(), false}
		addTask(task)
		newTable := table.New(os.Stdout)
		newTable.SetHeaders("ID", "Name", "Created")
		newTable.SetHeaderStyle(table.StyleBold)
		newTable.SetLineStyle(table.StyleBlue)
		newTable.SetDividers(table.UnicodeRoundedDividers)
		newTable.AddRow(strconv.Itoa(task.id), task.name, timediff.TimeDiff(task.created))
		newTable.Render()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
