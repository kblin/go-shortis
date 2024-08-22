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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <alias> <url>",
	Short: "Add a short URL alias",
	Long: `Add a short URL alias.

This adds a new URL alias to the link shortener.`,
	Args: cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := InitDb()
		if err != nil {
			panic(fmt.Errorf("error opening database: %s", err))
		}

		m := model.NewShortisModel(db)

		err = m.Add(args[0], args[1])
		if err != nil {
			panic(fmt.Errorf("error adding entry: %s", err))
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
