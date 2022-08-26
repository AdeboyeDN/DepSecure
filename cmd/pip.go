/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pipCmd represents the pip command
var pipCmd = &cobra.Command{
	Use:   "pip",
	Short: "Use to get pip dependencies informations",
	Long:  `Use to get pip dependencies informations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pip called")
	},
}

func init() {
	rootCmd.AddCommand(pipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
