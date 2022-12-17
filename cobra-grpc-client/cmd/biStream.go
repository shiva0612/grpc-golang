/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	helper "shiva/cobra-grpc-client/helper"

	"github.com/spf13/cobra"
)

// biStreamCmd represents the biStream command
var biStreamCmd = &cobra.Command{
	Use:   "biStream",
	Short: "A brief description of your command",
	Long:  `bi-directional streaming ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("biStream called")
		helper.Bistream()
	},
}

func init() {
	rootCmd.AddCommand(biStreamCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// biStreamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// biStreamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
