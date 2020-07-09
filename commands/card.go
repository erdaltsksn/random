package commands

import (
	"github.com/spf13/cobra"

	"github.com/erdaltsksn/random"
)

var cardType string

// cardCmd represents the card command.
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Generate a random card number",
	Long:  `Generates a random card number according to named card type.`,
	Run: func(cmd *cobra.Command, args []string) {
		printOutput(random.GetCardType(cardType).Generate())
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)

	// Here you will define your flags and configuration settings.
	cardCmd.Flags().StringVarP(&cardType, "type", "t", "",
		`Which card type should be used. Available cards:
- AmericanExpress
- DinersClub
- DinersClubUS
- Discover
- JCB
- Laser
- Maestro
- Mastercard
- Solo
- Unionpay
- Visa
- Mir
`)
}
