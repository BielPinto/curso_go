/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/BielPinto/curso_go/14-cobra_cli/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category`,
		RunE:  RunCreate(categoryDb),
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	cmd.Help()
	// },
	RunE: RunCreate(GetCategoryDB(GetDb())),
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	// createCmd.PersistentFlags().String("name", "", "name of the catgory")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RunCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		_, err := categoryDB.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}
