package main

import (
	"fmt"
	"os"

	"github.com/iMartyn/k8szoo/src"
	"github.com/spf13/cobra"
)

func main() {
	var animalname string

	var rootCmd = &cobra.Command{
		Use:   "k8szoo",
		Short: "k8szoo is a quick templating web frontend",
		Long: `k8szoo lets you template files (designed for yaml) with 
animal names and sounds`,
		Run: func(cmd *cobra.Command, args []string) {
			if (len(animalname) > 0) {
				animalFound := k8szoo.FindAnimal(animalname)
				if animalFound.AnimalName == "" {
					fmt.Printf("I don't know what sound a %s makes.\n", animalname)
				} else {
					fmt.Printf("The %s %s\n", animalFound.AnimalName, animalFound.AnimalSound)
				}
			} else {
				fmt.Println("Enter valid input. Hint, there isn't one!")
				cmd.Help()
			}
		},
	}
	var serveCmd = &cobra.Command{
		Use: "serve",
		Short: "Serve http requests",
		Long: "Run the webserver to serve http requests",
		Run: func(cmd *cobra.Command, args []string) {
			k8szoo.HandleHTTP()
		},
	}
	
	rootCmd.Flags().StringVarP(&animalname, "animalname", "p", "", "Animal name")
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
