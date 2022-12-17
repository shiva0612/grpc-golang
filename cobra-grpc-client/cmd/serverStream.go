/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	helper "shiva/cobra-grpc-client/helper"

	"github.com/spf13/cobra"
)

// serverStreamCmd represents the serverStream command
var serverStreamCmd = &cobra.Command{
	Use:   "serverStream",
	Short: "A brief description of your command",
	Long:  `server streaming ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serverStream called")
		helper.Sstream()
	},
}

func init() {
	rootCmd.AddCommand(serverStreamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverStreamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverStreamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
