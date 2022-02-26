/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/booscaaa/rtdd-golang/structs/adapter/sqlite"
	"github.com/booscaaa/rtdd-golang/structs/adapter/sqlite/person_repository"
	"github.com/booscaaa/rtdd-golang/structs/core/person_usecase"
	"github.com/spf13/cobra"
)

// personCmd represents the person command
var personCmd = &cobra.Command{
	Use:   "person",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := sqlite.GetConnection()
		sqlite.ConfigConnection(db)
		personRepository := person_repository.NewPersonRepository(db)
		personUsecase := person_usecase.NewPersonUseCase(personRepository)

		list, _ := cmd.Flags().GetBool("list")

		if list {
			people, _ := personUsecase.Fetch()
			fmt.Println(people)
		}
	},
}

func init() {
	rootCmd.AddCommand(personCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// personCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	personCmd.Flags().BoolP("list", "l", false, "Help message for toggle")
}
