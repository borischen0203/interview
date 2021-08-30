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
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// var questionsMap = make(map[int]string)

// askCmd represents the ask command
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Random an interview question",
	Long:  `This command print a Random an interview question`,
	Run: func(cmd *cobra.Command, args []string) {
		questionsMap, err := Load()
		if err != nil {
			log.Fatal(err)
		}
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(len(questionsMap)) + 1

		fmt.Println(questionsMap[num])
	},
}

//Load csv file with default interview questions
func Load() (map[int]string, error) {
	file, err := os.Open("default.csv")

	if err != nil {
		fmt.Println("An error encountered ::", err)
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	csvData, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	result := make(map[int]string)
	for num, line := range csvData {
		if num == 0 {
			continue
		}
		result[num] = line[1]
	}
	return result, nil
}

func init() {
	rootCmd.AddCommand(askCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// askCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
