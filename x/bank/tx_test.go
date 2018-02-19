package bank

import (
	"testing"

	"github.com/stretchr/testify/assert"

	crypto "github.com/tendermint/go-crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestInputValidation(t *testing.T) {
	addr1 := crypto.Address([]byte{1, 2})
	addr2 := crypto.Address([]byte{7, 8})
	someCoins := sdk.Coins{{"atom", 123}}
	multiCoins := sdk.Coins{{"atom", 123}, {"eth", 20}}

	var emptyAddr crypto.Address
	emptyCoins := sdk.Coins{}
	emptyCoins2 := sdk.Coins{{"eth", 0}}
	someEmptyCoins := sdk.Coins{{"eth", 10}, {"atom", 0}}
	minusCoins := sdk.Coins{{"eth", -34}}
	someMinusCoins := sdk.Coins{{"atom", 20}, {"eth", -34}}
	unsortedCoins := sdk.Coins{{"eth", 1}, {"atom", 1}}

	cases := []struct {
		valid bool
		txIn  Input
	}{
		// auth works with different apps
		{true, NewInput(addr1, someCoins)},
		{true, NewInput(addr2, someCoins)},
		{true, NewInput(addr2, multiCoins)},

		{false, NewInput(emptyAddr, someCoins)},  // empty address
		{false, NewInput(addr1, emptyCoins)},     // invalid coins
		{false, NewInput(addr1, emptyCoins2)},    // invalid coins
		{false, NewInput(addr1, someEmptyCoins)}, // invalid coins
		{false, NewInput(addr1, minusCoins)},     // negative coins
		{false, NewInput(addr1, someMinusCoins)}, // negative coins
		{false, NewInput(addr1, unsortedCoins)},  // unsorted coins
	}

	for i, tc := range cases {
		err := tc.txIn.ValidateBasic()
		if tc.valid {
			assert.Nil(t, err, "%d: %+v", i, err)
		} else {
			assert.NotNil(t, err, "%d", i)
		}
	}
}

func TestOutputValidation(t *testing.T) {
	addr1 := crypto.Address([]byte{1, 2})
	addr2 := crypto.Address([]byte{7, 8})
	someCoins := sdk.Coins{{"atom", 123}}
	multiCoins := sdk.Coins{{"atom", 123}, {"eth", 20}}

	var emptyAddr crypto.Address
	emptyCoins := sdk.Coins{}
	emptyCoins2 := sdk.Coins{{"eth", 0}}
	someEmptyCoins := sdk.Coins{{"eth", 10}, {"atom", 0}}
	minusCoins := sdk.Coins{{"eth", -34}}
	someMinusCoins := sdk.Coins{{"atom", 20}, {"eth", -34}}
	unsortedCoins := sdk.Coins{{"eth", 1}, {"atom", 1}}

	cases := []struct {
		valid bool
		txOut Output
	}{
		// auth works with different apps
		{true, NewOutput(addr1, someCoins)},
		{true, NewOutput(addr2, someCoins)},
		{true, NewOutput(addr2, multiCoins)},

		{false, NewOutput(emptyAddr, someCoins)},  // empty address
		{false, NewOutput(addr1, emptyCoins)},     // invalid coins
		{false, NewOutput(addr1, emptyCoins2)},    // invalid coins
		{false, NewOutput(addr1, someEmptyCoins)}, // invalid coins
		{false, NewOutput(addr1, minusCoins)},     // negative coins
		{false, NewOutput(addr1, someMinusCoins)}, // negative coins
		{false, NewOutput(addr1, unsortedCoins)},  // unsorted coins
	}

	for i, tc := range cases {
		err := tc.txOut.ValidateBasic()
		if tc.valid {
			assert.Nil(t, err, "%d: %+v", i, err)
		} else {
			assert.NotNil(t, err, "%d", i)
		}
	}
}

func TestSendMsgValidation(t *testing.T) {

	addr1 := crypto.Address([]byte{1, 2})
	addr2 := crypto.Address([]byte{7, 8})
	atom123 := sdk.Coins{{"atom", 123}}
	atom124 := sdk.Coins{{"atom", 124}}
	eth123 := sdk.Coins{{"eth", 123}}
	atom123eth123 := sdk.Coins{{"atom", 123}, {"eth", 123}}

	input1 := NewInput(addr1, atom123)
	input2 := NewInput(addr1, eth123)
	output1 := NewOutput(addr2, atom123)
	output2 := NewOutput(addr2, atom124)
	output3 := NewOutput(addr2, eth123)
	outputMulti := NewOutput(addr2, atom123eth123)

	var emptyAddr crypto.Address

	cases := []struct {
		valid bool
		tx    SendMsg
	}{
		{false, SendMsg{}},                           // no input or output
		{false, SendMsg{Inputs: []Input{input1}}},    // just input
		{false, SendMsg{Outputs: []Output{output1}}}, // just ouput
		{false, SendMsg{
			Inputs:  []Input{NewInput(emptyAddr, atom123)}, // invalid input
			Outputs: []Output{output1}}},
		{false, SendMsg{
			Inputs:  []Input{input1},
			Outputs: []Output{{emptyAddr, atom123}}}, // invalid ouput
		},
		{false, SendMsg{
			Inputs:  []Input{input1},
			Outputs: []Output{output2}}, // amounts dont match
		},
		{false, SendMsg{
			Inputs:  []Input{input1},
			Outputs: []Output{output3}}, // amounts dont match
		},
		{false, SendMsg{
			Inputs:  []Input{input1},
			Outputs: []Output{outputMulti}}, // amounts dont match
		},
		{false, SendMsg{
			Inputs:  []Input{input2},
			Outputs: []Output{output1}}, // amounts dont match
		},

		{true, SendMsg{
			Inputs:  []Input{input1},
			Outputs: []Output{output1}},
		},
		{true, SendMsg{
			Inputs:  []Input{input1, input2},
			Outputs: []Output{outputMulti}},
		},
	}

	for i, tc := range cases {
		err := tc.tx.ValidateBasic()
		if tc.valid {
			assert.Nil(t, err, "%d: %+v", i, err)
		} else {
			assert.NotNil(t, err, "%d", i)
		}
	}
}

/*
// TODO where does this test belong ?
func TestSendMsgSigners(t *testing.T) {
	signers := []crypto.Address{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	someCoins := sdk.Coins{{"atom", 123}}
	inputs := make([]Input, len(signers))
	for i, signer := range signers {
		inputs[i] = NewInput(signer, someCoins)
	}
	tx := NewSendMsg(inputs, nil)

	assert.Equal(t, signers, tx.Signers())
}
*/
