/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// containersCmd represents the containers command
var containersCmd = &cobra.Command{
	Use:   "containers",
	Short: "Keyfactor CertificateStoreContainer API and utilities.",
	Long:  `A collections of APIs and utilities for interacting with Keyfactor certificate store containers.`,
}

var containersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create certificate store container.",
	Long:  `Create certificate store container.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create store containers not implemented.")
	},
}

var containersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certificate store container by ID or name.",
	Long:  `Get certificate store container by ID or name.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetOutput(ioutil.Discard)
		client := cmd.Flag("client").Value.String()
		kfClient, _ := initClient()
		agents, aErr := kfClient.GetStoreContainer(client)
		if aErr != nil {
			fmt.Printf("Error, unable to get orchestrator %s. %s\n", client, aErr)
			log.Fatalf("Error: %s", aErr)
		}
		output, jErr := json.Marshal(agents)
		if jErr != nil {
			fmt.Printf("Error invalid API response from Keyfactor. %s\n", jErr)
			log.Fatalf("[ERROR]: %s", jErr)
		}
		fmt.Printf("%s", output)
	},
}

var containersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update certificate store container by ID or name.",
	Long:  `Update certificate store container by ID or name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Update store containers not implemented.")
	},
}

var containersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete certificate store container by ID or name.",
	Long:  `Delete certificate store container by ID or name.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete store containers not implemented.")
	},
}

var containersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List certificate store containers.",
	Long:  `List certificate store containers.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetOutput(ioutil.Discard)
		kfClient, _ := initClient()
		agents, aErr := kfClient.GetStoreContainers()
		if aErr != nil {
			fmt.Printf("Error, unable to list store containers. %s\n", aErr)
			log.Fatalf("Error: %s", aErr)
		}
		output, jErr := json.Marshal(agents)
		if jErr != nil {
			fmt.Printf("Error invalid API response from Keyfactor. %s\n", jErr)
			log.Fatalf("[ERROR]: %s", jErr)
		}
		fmt.Printf("%s", output)
	},
}

func init() {
	rootCmd.AddCommand(containersCmd)
	// LIST containers command
	containersCmd.AddCommand(containersListCmd)
	// GET containers command
	containersCmd.AddCommand(containersGetCmd)
	containersGetCmd.Flags().StringP("id", "i", "", "ID or name of the cert store container.")
	// CREATE containers command
	//containersCmd.AddCommand(containersCreateCmd)
	// UPDATE containers command
	//containersCmd.AddCommand(containersUpdateCmd)
	// DELETE containers command
	//containersCmd.AddCommand(containersDeleteCmd)
	// Utility functions
}
