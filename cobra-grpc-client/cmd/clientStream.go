/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	helper "shiva/cobra-grpc-client/helper"

	"github.com/spf13/cobra"
)

// clientStreamCmd represents the clientStream command
var clientStreamCmd = &cobra.Command{
	Use:   "clientStream",
	Short: "A brief description of your command",
	Long:  `client streaming ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clientStream called")
		helper.Cstream()
	},
}

func init() {
	rootCmd.AddCommand(clientStreamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientStreamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientStreamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
