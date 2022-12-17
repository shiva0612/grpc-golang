/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"shiva/cobra-grpc-client/cmd"
	helper "shiva/cobra-grpc-client/helper"
	us "shiva/proto/services/user"

	"google.golang.org/grpc"
)

func init() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	// defer cc.Close()

	helper.Client = us.NewUserApiClient(cc)

}
func main() {
	cmd.Execute()
}
