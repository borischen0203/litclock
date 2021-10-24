package cmd

/***
 *
 * This file mainly handle root command
 *
 * @author: Boris
 * @version: 2021-10-22
 *
 */

import (
	"fmt"
	"os"
	"time"

	"github.com/borischen0203/litclock/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// //go:embed gophers
// var embedGopherFiles embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "litclock",
	Short: "This command convert numeric time to human friendly text",
	Long:  `This command convert numeric time to human friendly text`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: NumericToTextCommand,
}

//This function mainly excute root command
func NumericToTextCommand(cmd *cobra.Command, args []string) {
	result := NumericToText(args...)
	if result == "Invalid input" {
		services.GopherSay(result)
		// fmt.Println(result)
	} else {
		fmt.Println(result)
	}
}

//This function mainly excute time convert function
func NumericToText(args ...string) string {
	inputLength := len(args)
	humanText := ""

	switch inputLength {
	case 0: //without input paramater
		t := time.Now()
		humanText = services.TimeToWords(t.Hour(), t.Minute())
	case 1: //with input one paramater
		t, err := time.Parse("15:04", args[0])
		if err != nil {
			return "Invalid input"
		}
		humanText = services.TimeToWords(t.Hour(), t.Minute())
	default: //with input more than one paramater
		return "Invalid input"

	}
	return humanText
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.litclock.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".litclock" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".litclock")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
