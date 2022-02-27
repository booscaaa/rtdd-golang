/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	personservice "github.com/booscaaa/rtdd-golang/gateway/adapter/grpcservice/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:1111", grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()

		productServ := personservice.NewPersonServiceClient(conn)

		response, err := productServ.Fetch(context.Background(), &personservice.FetchRequest{})
		if err != nil {
			fmt.Println(err)
		} else {
			people := response.People

			for _, person := range people {
				fmt.Println("Person Info")
				fmt.Println("id: ", person.Id)
				fmt.Println("name: ", person.Name)
				fmt.Println("age: ", person.Age)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}