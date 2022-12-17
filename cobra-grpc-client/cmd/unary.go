/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	helper "shiva/cobra-grpc-client/helper"

	"github.com/spf13/cobra"
)

// unaryCmd represents the unary command
var unaryCmd = &cobra.Command{
	Use:   "unary",
	Short: "A brief description of your command",
	Long:  `unary call ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unary called")
		helper.Unary()
	},
}

func init() {
	rootCmd.AddCommand(unaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
