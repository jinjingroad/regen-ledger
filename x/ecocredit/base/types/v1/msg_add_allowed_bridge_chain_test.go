package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddAllowedBridgeChain struct {
	t         gocuke.TestingT
	msg       *MsgAddAllowedBridgeChain
	signBytes string
	err       error
}

func TestMsgAddAllowedBridgeChain(t *testing.T) {
	gocuke.NewRunner(t, &msgAddAllowedBridgeChain{}).Path("./features/msg_add_allowed_bridge_chain.feature").Run()
}

func (s *msgAddAllowedBridgeChain) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddAllowedBridgeChain) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAddAllowedBridgeChain{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddAllowedBridgeChain) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddAllowedBridgeChain) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAddAllowedBridgeChain) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAddAllowedBridgeChain) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgAddAllowedBridgeChain) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
