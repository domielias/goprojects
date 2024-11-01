/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "This command allow to show undone tasks",
	Long:  "This command allow to show all undone tasks",
	Run: func(cmd *cobra.Command, args []string) {
		rawTasks, err := getFileData()
		if err != nil {
			log.Fatal(err)
		}
		tasks, err := getListTasks(rawTasks)
		if err != nil {
			log.Fatal(err)
		}
		newTable := table.New(os.Stdout)
		newTable.SetHeaders("ID", "Name", "Created")
		newTable.SetHeaderStyle(table.StyleBold)
		newTable.SetLineStyle(table.StyleBlue)
		newTable.SetDividers(table.UnicodeRoundedDividers)
		if len(tasks) < 1 {
			tml.Printf("<white>We don't have any pending task</white>\n")
			return
		}
		for _, task := range tasks {
			if task.isComplete == true {
				continue
			}
			newTable.AddRow(strconv.Itoa(task.id), task.name, timediff.TimeDiff(task.created))
		}
		newTable.Render()
	},
}

var listCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "This command allow to show complete tasks",
	Long:  "This command allow to show all complete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		rawTasks, err := getFileData()
		if err != nil {
			log.Fatal(err)
		}
		tasks, err := getListTasks(rawTasks)
		if err != nil {
			log.Fatal(err)
		}
		newTable := table.New(os.Stdout)
		newTable.SetHeaders("ID", "Name", "Created")
		newTable.SetHeaderStyle(table.StyleBold)
		newTable.SetLineStyle(table.StyleBlue)
		newTable.SetDividers(table.UnicodeRoundedDividers)
		if len(tasks) < 1 {
			tml.Printf("<white>We don't have any pending task</white>\n")
			return
		}
		for _, task := range tasks {
			if task.isComplete == true {
				newTable.AddRow(strconv.Itoa(task.id), task.name, timediff.TimeDiff(task.created))
			}
		}
		newTable.Render()
	},
}

var listAllCmd = &cobra.Command{
	Use:   "all",
	Short: "This command allow to show all tasks",
	Long:  "This command allow to show all tasks (undone and done tasks)",

	Run: func(cmd *cobra.Command, args []string) {
		rawTasks, err := getFileData()
		if err != nil {
			log.Fatal(err)
		}
		tasks, err := getListTasks(rawTasks)
		if err != nil {
			log.Fatal(err)
		}
		newTable := table.New(os.Stdout)
		newTable.SetHeaders("ID", "Name", "Created")
		newTable.SetHeaderStyle(table.StyleBold)
		newTable.SetLineStyle(table.StyleBlue)
		newTable.SetDividers(table.UnicodeRoundedDividers)
		if len(tasks) < 1 {
			tml.Printf("<white>We don't have any pending task</white>\n")
			return
		}
		for _, task := range tasks {
			newTable.AddRow(strconv.Itoa(task.id), task.name, timediff.TimeDiff(task.created))
		}
		newTable.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listCompleteCmd)
	listCmd.AddCommand(listAllCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
