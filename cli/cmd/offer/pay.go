// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package offer

import (
	"time"

	"github.com/lightsparkdev/go-sdk/cli/sdk"
	"github.com/lightsparkdev/go-sdk/objects"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// payCmd represents the pay command
var payCmd = &cobra.Command{
	Use:   "pay",
	Short: "Pay a Bolt12 offer.",
	Long: `
Pay a Bolt12 offer. For example:

    lightspark-cli offer pay --node-id <node-id> --offer <lno...> --timeout 30s --maxfee 1000

If the offer does not have a minimum amount, then the amount must be specified.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		nodeId := viper.GetString("LIGHTSPARK_NODE_ID")

		offer, err := cmd.Flags().GetString("offer")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		timeout, err := cmd.Flags().GetDuration("timeout")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		maxFeeMsats, err := cmd.Flags().GetInt64("maxfee")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		amountFlag, err := cmd.Flags().GetInt64("amount")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		var amountMsats *int64

		if amountFlag > 0 {
			amountMsats = &amountFlag
		}

		wait, err := cmd.Flags().GetBool("wait")
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		client, err := sdk.GetClientWithSigningKey()
		if err != nil {
			cmd.PrintErrf("Failed to get client with signing key: %s\n", err)
			return
		}

		payment, err := client.PayOffer(nodeId, offer, int(timeout.Seconds()), maxFeeMsats, amountMsats, nil)
		if err != nil {
			cmd.PrintErrf("Failed to pay offer: %s\n", err)
			return
		}

		if !wait {
			cmd.Println(payment.Id)
		} else {
			paymentId := payment.Id
			startTime := time.Now()

			cmd.Printf("Waiting for payment %s to complete...\n", paymentId)

			for payment.GetStatus() != objects.TransactionStatusSuccess && payment.GetStatus() != objects.TransactionStatusFailed {
				cmd.Printf("Payment status: %s, Sleeping for 5 seconds before refetching...\n", payment.GetStatus().StringValue())
				time.Sleep(5 * time.Second)
				if time.Since(startTime) > time.Minute*2 {
					cmd.PrintErrf("Payment wasn't updated after two minutes, giving up...\n")
					return
				}

				entity, err := client.GetEntity(paymentId)
				if err != nil {
					cmd.PrintErrf("Failed to get entity: %s\n", err)
					return
				}

				castPayment, didCast := (*entity).(objects.OutgoingPayment)
				if !didCast {
					cmd.PrintErrf("Failed to cast payment to Transaction: %s\n", err)
					return
				}

				payment = &castPayment
			}

			if payment.GetStatus() == objects.TransactionStatusFailed {
				cmd.PrintErrf("Payment failed: %s\n", payment.FailureReason.StringValue())
			} else {
				cmd.Println("Payment successful!")
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// payCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// payCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	payCmd.Flags().String("offer", "", "The bech32 encoded Bolt12 offer to pay.")
	payCmd.MarkFlagRequired("offer")

	payCmd.Flags().Int64("maxfee", 0, "The maximum fee in fees to pay for this payment, in millisatoshis.")
	payCmd.MarkFlagRequired("maxfee")

	payCmd.Flags().Duration("timeout", 30*time.Second, "The timeout during which the payment will be attempted.")
	payCmd.Flags().Int64("amount", 0, "The amount in millisatoshis to pay for this offer. Only required if the offer does not have a minimum amount.")
	payCmd.Flags().Bool("wait", false, "Wait for the payment to complete.")
}
