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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an alias from the database",
	Long:  `Remove an alias from the database.`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := InitDb()
		if err != nil {
			panic(fmt.Errorf("error opening database: %s", err))
		}

		m := model.NewShortisModel(db)

		err = m.Remove(args[0])
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				os.Exit(0)
			}
			panic(fmt.Errorf("Error removing from database: %s", err))
		}

		fmt.Println("Removed", args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
