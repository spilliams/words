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

var length, minLength, maxLength int
var fill, set string

func init() {
	RootCmd.AddCommand(clueCmd)

	clueCmd.Flags().IntVarP(&length, "length", "l", -1, "The length of the desired word.")
	clueCmd.Flags().IntVarP(&minLength, "min", "", -1, "The minimum length of the desired word.")
	clueCmd.Flags().IntVarP(&maxLength, "max", "", -1, "The maximum length of the desired word.")
	clueCmd.Flags().StringVarP(&fill, "fill", "f", "", "The partially-filled in word, using spaces for unknown letters.")
	clueCmd.Flags().StringVarP(&set, "set", "s", "", "The set of letters available. Note: they won't repeat usage.")
}

func clue(cmd *cobra.Command, args []string) error {
	if length == -1 && len(fill) == 0 && len(set) == 0 {
		return fmt.Errorf("You must provide a value for at least one of: length, fill, set")
	}

	count := 0
	for _, word := range util.Words {

		ok := true

		// if word length correct
		if length != -1 && len(word) != length {
			ok = false
		}
		if minLength != -1 && len(word) < minLength {
			ok = false
		}
		if maxLength != -1 && len(word) > maxLength {
			ok = false
		}

		// if word match fill
		if len(fill) != 0 {
			if !util.WordMatchesFill(word, fill) {
				ok = false
			}
		}

		// if word match set
		if len(set) != 0 {
			if !util.WordMatchesSet(word, set) {
				ok = false
			}
		}

		if ok {
			count++
			fmt.Println(word)
		}
	}
	fmt.Printf("%v matching words\n", count)
	return nil
}
