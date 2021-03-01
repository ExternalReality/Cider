/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	teamcity2 "github.com/cvbarros/go-teamcity/teamcity"
	"github.com/spf13/cobra"
	"net/http"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "It allows to enable provider",
	Long: `this command is used to enable provider by their name
          e.g. cider enable provider teamcity /cider enable provider jenkins.`,
	Run: func(cmd *cobra.Command, args []string) {
		enableProv(args)
	},
}

func enableProv(args []string) {
	if len(args) ==2{
		if args[0]== "provider" && args[1]=="teamcity"{
			client, err:= teamcity2.NewClientWithAddress(teamcity2.BasicAuth("admin","admin"),"http://localhost:8081",http.DefaultClient)
			if err!=nil{
				fmt.Println(err.Error())
			}
			result,err1:=client.Validate()
			if err1!=nil{
				fmt.Println(err1.Error())
			}else{
				fmt.Println("Task status :",result)

			}

		}else{
			fmt.Print("Enter valid command  e.g. cider enable provider teamcity")
		}
	}
}

func init() {
	rootCmd.AddCommand(enableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
