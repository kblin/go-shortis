/*
Copyright © 2024 Kai Blin

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

	"github.com/spf13/cobra"

	"kblin.org/shortis/internal/model"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all current alias entries",
	Long:  `List all current alias entries.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := InitDb()
		if err != nil {
			panic(fmt.Errorf("error opening database: %s", err))
		}

		m := model.NewShortisModel(db)

		links, err := m.List()
		if err != nil {
			panic(fmt.Errorf("error listing entries: %s", err))
		}

		fmt.Println("#\tAlias\tURL")
		for i, link := range links {
			fmt.Printf("%d\t%s\t%s\n", i+1, link.Short, link.Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
