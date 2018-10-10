// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/suburban/flexi-pass/enviro"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

// userAdd represents the user command
var userAdd = &cobra.Command{
	Use:   "user-add",
	Short: "Add a new flexi-pass user",
	Long:  `Create a new flexi-pass user in the system.`,
	Run:   runCommand,
}

// Setup a series of questions to get the user's details.
var userQuestions = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "Name"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "email",
		Prompt:    &survey.Input{Message: "Email"},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
}

func init() {
	rootCmd.AddCommand(userAdd)
}

func runCommand(cmd *cobra.Command, args []string) {
	log.WithFields(log.Fields{
		"environment": enviro.Env,
	}).Debug("Environment setup")

	// answers := struct {
	// 	Name  string
	// 	Email string
	// }{}

	// // perform the questions
	// err := survey.Ask(userQuestions, &answers)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Printf("%s, %s.", answers.Name, answers.Email)
}
