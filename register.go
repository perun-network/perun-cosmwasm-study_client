package main

import (
	"encoding/json"

	types "github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type RegisterMsg struct {
	Params     ChannelParameters `json:"params"`
	State      ChannelState      `json:"state"`
	Signatures [2]Signature      `json:"sigs"`
}

func (c *ChannelClient) register(ch *Channel, sig1 Signature, sig2 Signature) {
	msg, err := genRegisterMsg(
		c.ctx.FromAddress,
		c.contractAddress,
		ch,
		sig1,
		sig2,
	)
	if err != nil {
		panic(err)
	}

	validateMessageJSON(msg.Msg)

	_, err = transact(c.ctx, &msg)
	if err != nil {
		panic(err)
	}
}

func genRegisterMsg(
	sender sdk.AccAddress,
	contract ContractAddress,
	ch *Channel,
	sig1 Signature,
	sig2 Signature,
) (msg types.MsgExecuteContract, err error) {
	_msg, err := json.Marshal(
		map[string]interface{}{
			"register": RegisterMsg{
				Params:     ch.params,
				State:      ch.state,
				Signatures: [2]Signature{sig1, sig2},
			},
		},
	)
	if err != nil {
		return
	}

	msg = types.MsgExecuteContract{
		Sender:   sender.String(),
		Contract: contract,
		Msg:      _msg,
		Funds:    nil,
	}
	return
}
