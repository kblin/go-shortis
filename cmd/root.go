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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
)

var (
	cfgFile string
	dbFile  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shortis",
	Short: "A small stand-alone URL shortener",
	Long: `A small stand-alone URL shortener.

	Now that more and more URL shorteners are shutting down, it's time to
	host your own.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbfile := viper.GetString("db")
		fmt.Println("Using database: ", dbfile)
	},
	Version: "0.1.0",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(readConfig)
	viper.SetEnvPrefix("shortis")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/shortis.toml)")
	viper.BindEnv("config")
	rootCmd.PersistentFlags().StringVar(&dbFile, "db", "shortis.db", "SQLite database file")
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("db"))
	viper.BindEnv("db")

}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite", viper.GetString("db"))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func readConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/shortis")
		viper.SetConfigName("shortis")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading configuration:", err)
		os.Exit(1)
	}
}
