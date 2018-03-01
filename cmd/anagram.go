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
