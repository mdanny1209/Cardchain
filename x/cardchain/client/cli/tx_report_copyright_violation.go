package cli

import (
	"strconv"

	//"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdReportCopyrightViolation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report-copyright-violation [card-id] [deposit] [original-artist] [link] [additional-info]",
		Short: "Broadcast message reportCopyrightViolation",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCardId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argDeposit := args[1]
			argOriginalArtist := args[2]
			argLink := args[3]
			argAdditionalInfo := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(argDeposit)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress()
			cardId := strconv.FormatUint(argCardId, 10)

			content := gov.NewTextProposal(
				"Copyright violation on card" + cardId,
				"The Card `" + cardId + "` is a copyright violation.\n" + argAdditionalInfo + "\n\nOriginal artist: `" + argOriginalArtist + "`\nLink: " + argLink,
			)

			msg, err := gov.NewMsgSubmitProposal(
				content,
				deposit,
				creator,
			)
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
