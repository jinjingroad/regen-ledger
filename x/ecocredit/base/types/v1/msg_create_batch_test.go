package v1

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateBatch struct {
	t         gocuke.TestingT
	msg       *MsgCreateBatch
	err       error
	signBytes string
}

func TestMsgCreateBatch(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateBatch{}).Path("./features/msg_create_batch.feature").Run()
}

func (s *msgCreateBatch) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateBatch) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateBatch{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateBatch) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgCreateBatch) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateBatch) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateBatch) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateBatch) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgCreateBatch) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
