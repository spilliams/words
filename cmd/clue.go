package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/words/util"
)

// clueCmd represents the clue command
var clueCmd = &cobra.Command{
	Use:   "clue",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: clue,
}

var length int
var fill, set string

func init() {
	RootCmd.AddCommand(clueCmd)

	clueCmd.Flags().IntVarP(&length, "length", "l", -1, "The length of the desired word.")
	clueCmd.Flags().StringVarP(&fill, "fill", "f", "", "The partially-filled in word, using spaces for unknown letters.")
	clueCmd.Flags().StringVarP(&set, "set", "s", "", "The set of letters available. Note: they won't repeat usage.")
}

func clue(cmd *cobra.Command, args []string) error {
	fmt.Println("clue command called")
	if length == -1 && len(fill) == 0 && len(set) == 0 {
		return fmt.Errorf("You must provide a value for at least one of: length, fill, set")
	}

	count := 0
	for _, word := range util.Words {

		// if word length correct
		ok := true
		if length != -1 && len(word) != length {
			ok = false
		}

		// if word match fill
		if len(fill) != 0 {
			if len(word) != len(fill) {
				ok = false
			} else {
				for i := 0; i < len(fill); i++ {
					if fill[i:i+1] != " " && word[i:i+1] != fill[i:i+1] {
						ok = false
					}
				}
			}

		}

		// if word match set
		// TODO

		if ok {
			count++
			fmt.Println(word)
		}
	}
	fmt.Printf("%v matching words\n", count)
	return nil
}
