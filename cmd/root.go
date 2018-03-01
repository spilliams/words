package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spilliams/words/util"
)

var cfgFile, dictionary string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "words",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.words.yaml)")
	RootCmd.PersistentFlags().StringVarP(&dictionary, "dictionary", "d", "", "dictionary file to use. Separate words with line-breaks.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".words" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".words")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Words returns the set of words this command uses
func Words() []string {
	words := util.Words()

	// load the dictionary file, if given
	if len(dictionary) != 0 {
		bytes, err := ioutil.ReadFile(dictionary)
		if err != nil {
			fmt.Printf("Error reading dictionary file %v: %v\n", dictionary, err)
			fmt.Println("Using included dictionary instead")
			return words
		}
		words = strings.Split(string(bytes), "\n")
	}

	// make sure there are no blanks or dupes
	words = append(words, "")
	words = util.ArrayUnique(words)
	words = words[:len(words)-1]

	fmt.Printf("Loaded dictionary of %v words\n", len(words))
	return words
}
