package main

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	junomessages "github.com/forbole/juno/v3/modules/messages"

	"github.com/AutonomyNetwork/autonomy-chain/x/issuance/types"
	nfttypes "github.com/AutonomyNetwork/nft/types"
)

var autonomyMessageAddressesParser = junomessages.JoinMessageParsers(
	issanceMessageAddressesParser,
	nftMessageAddressesParser,
)

func issanceMessageAddressesParser(_ codec.Codec, cosmosmsg sdk.Msg) ([]string, error) {
	switch msg := cosmosmsg.(type) {
	case *types.MsgIssueToken:
		return []string{msg.Creator}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosmsg)
}
func nftMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {
	case *nfttypes.MsgCreateDenom:
		return []string{msg.Creator}, nil
	case *nfttypes.MsgMintNFT:
		return []string{msg.Creator}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)

}
