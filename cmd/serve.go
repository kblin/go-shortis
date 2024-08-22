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
	"github.com/spf13/viper"

	"kblin.org/shortis/internal/model"
	"kblin.org/shortis/internal/web"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the shortis web interface",
	Long: `Run the shortis web interface.

	This runs the web UI and API for shortis.`,
	Run: func(cmd *cobra.Command, args []string) {
		address := viper.GetString("address")
		port := viper.GetInt("port")
		db, err := InitDb()
		if err != nil {
			panic(fmt.Errorf("error opening database: %s", err))
		}

		m := model.NewShortisModel(db)

		web.Run(address, port, &m)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().Int("port", 7647, "Port to run the URL shortening server on")
	viper.BindEnv("port")
	viper.BindPFlag("port", serveCmd.PersistentFlags().Lookup("port"))

	serveCmd.PersistentFlags().String("address", "", "Server address to use")
	viper.BindEnv("address")
	viper.BindPFlag("address", serveCmd.PersistentFlags().Lookup("address"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
