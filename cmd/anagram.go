package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/words/util"
)

// anagramCmd represents the anagram command
var anagramCmd = &cobra.Command{
	Use:   "anagram WORD",
	Short: "Lists anagrams of a given word",
	Long:  `Lists anagrams of a given word`,
	RunE:  anagram,
}

func init() {
	RootCmd.AddCommand(anagramCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// anagramCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// anagramCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func anagram(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("anagram requires 1 argument")
	}
	count := 0
	for _, word := range Words() {

		if len(word) != len(args[0]) {
			continue
		}

		if word == args[0] {
			continue
		}

		if util.NormalizeString(word) != util.NormalizeString(args[0]) {
			continue
		}

		count++
		fmt.Println(word)
	}
	fmt.Printf("%v anagrams\n", count)
	return nil
}
