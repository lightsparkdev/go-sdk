// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package offer

import (
	"github.com/lightsparkdev/go-sdk/cli/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/terminal"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		nodeId := viper.GetString("LIGHTSPARK_NODE_ID")

		amountFlag, err := cmd.Flags().GetInt64("amount")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		var amountMsats *int64

		if amountFlag > 0 {
			amountMsats = &amountFlag
		}

		descriptionFlag, err := cmd.Flags().GetString("description")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		var description *string
		if descriptionFlag != "" {
			description = &descriptionFlag
		}

		client, err := sdk.GetClientWithSigningKey()
		if err != nil {
			cmd.PrintErrf("Failed to get client with signing key: %s\n", err)
			return
		}

		offer, err := client.CreateOffer(nodeId, amountMsats, description)
		if err != nil {
			cmd.PrintErrf("Failed to create offer: %s\n", err)
			return
		}

		cmd.Println(offer.EncodedOffer)

		qrc, _ := qrcode.New(offer.EncodedOffer)
		w := terminal.New()

		if err := qrc.Save(w); err != nil {
			cmd.PrintErrf("Failed to write QR code: %s\n", err)
			return
		}
	},
}

func init() {
	OfferCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().Int64("amount", 0, "The amount in millisatoshis required to pay for this offer.")
	createCmd.Flags().String("description", "", "A brief description of the offer.")
}
