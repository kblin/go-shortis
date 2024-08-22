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
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"kblin.org/shortis/internal/model"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <alias>",
	Short: "Get the URL for a specific alias",
	Long:  `Get the URL for a specific alias.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := InitDb()
		if err != nil {
			panic(fmt.Errorf("error opening database: %s", err))
		}

		m := model.NewShortisModel(db)

		url, err := m.Get(args[0])
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				fmt.Fprintln(os.Stderr, "No URL set up for", args[0])
				os.Exit(1)
			}
			panic(fmt.Errorf("error getting URL from database: %s", err))
		}
		fmt.Println(url)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
