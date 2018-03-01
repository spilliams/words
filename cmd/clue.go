package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spilliams/words/util"
)

// clueCmd represents the clue command
var clueCmd = &cobra.Command{
	Use:   "clue",
	Short: "Lists words matching a given clue.",
	Long: `Lists words matching a given clue.

A clue can contain any information about

- length
- the set of letters available to make the word
- the existing "fill", or positions of known letters in the word`,
	RunE: clue,
}

var length, minLength, maxLength int
var fill, set, setReuse string

func init() {
	RootCmd.AddCommand(clueCmd)

	clueCmd.Flags().IntVarP(&length, "length", "l", -1, "The length of the desired word.")
	clueCmd.Flags().IntVarP(&minLength, "min", "", -1, "The minimum length of the desired word.")
	clueCmd.Flags().IntVarP(&maxLength, "max", "", -1, "The maximum length of the desired word.")
	clueCmd.Flags().StringVarP(&fill, "fill", "f", "", "The partially-filled in word, using spaces for unknown letters.")
	clueCmd.Flags().StringVarP(&set, "set", "s", "", "The set of letters available. Letters are not reused (e.g. the set \"tis\" will match \"its\" but not \"sits\")")
	clueCmd.Flags().StringVarP(&setReuse, "set-reuse", "r", "", "The set of letters avaialble. Letters are reused (e.g. the set \"tis\" will match \"its\" and \"sits\")")
}

func clue(cmd *cobra.Command, args []string) error {
	if length == -1 && len(fill) == 0 && len(set) == 0 && len(setReuse) == 0 {
		return fmt.Errorf("You must provide a value for at least one of: length, fill, set, set-reuse")
	}

	count := 0
	for _, word := range Words() {

		// if word length correct
		if length != -1 && len(word) != length {
			continue
		}
		if minLength != -1 && len(word) < minLength {
			continue
		}
		if maxLength != -1 && len(word) > maxLength {
			continue
		}

		// if word match fill
		if len(fill) != 0 {
			if !util.WordMatchesFill(word, fill) {
				continue
			}
		}

		// if word match set
		if len(set) != 0 {
			if !util.WordMatchesSet(word, set, false) {
				continue
			}
		}
		if len(setReuse) != 0 {
			if !util.WordMatchesSet(word, setReuse, true) {
				continue
			}
		}

		count++
		fmt.Println(word)
	}
	fmt.Printf("%v matching words\n", count)
	return nil
}
