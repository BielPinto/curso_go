/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := cmd.Flags().GetString("name")
		// fmt.Println("category called with name: " + category)
		// exists, _ := cmd.Flags().GetBool("exists")
		// fmt.Println("Category called with exists: " + fmt.Sprint(exists))
		// id, _ := cmd.Flags().GetInt16("id")
		// fmt.Println("category called with id: ", fmt.Sprint(id))
		db := GetDb()
		category := GetCategoryDB(db)
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		category.Create(name, description)
		cmd.Help()
	},
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Called before of Run")
	// },
	// PostRun: func(cmd *cobra.Command, agrs []string) {
	// 	fmt.Println("Called after of Run")
	// },
	// RunE: func(cmd *cobra.Command, agrs []string) error {
	// 	return fmt.Errorf("called when ocurrent error ")
	// },
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("Descripiton", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")

	// ategoryCmd.Flags().String("name", "", "name of the category")
	// categoryCmd.PersistentFlags().String("name", "", "name of the category")
	// categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Category name")
	// categoryCmd.PersistentFlags().StringP("name", "n", "Y", "name of the catgeory")       //go run main.go category -n=x
	// categoryCmd.PersistentFlags().BoolP("exists", "e", false, "check if category exists") // go run main.go category -e=false
	// categoryCmd.PersistentFlags().Int16P("id", "i", 0, "Id of the category")              // go run main.go category -n=categoria -e --id=10
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
