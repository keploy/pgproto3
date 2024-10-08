package pgproto3

import (
	"bytes"
	"errors"

	"github.com/jackc/pgio"
)

type SASLInitialResponse struct {
	AuthMechanism string `json:"auth_mechanism" yaml:"auth_mechanism"`
	Data          []byte `json:"data" yaml:"data"`
}

// Frontend identifies this message as sendable by a PostgreSQL frontend.
func (*SASLInitialResponse) Frontend() {}

// Decode decodes src into dst. src must contain the complete message with the exception of the initial 1 byte message
// type identifier and 4 byte message length.
func (dst *SASLInitialResponse) Decode(src []byte) error {
	//println("SASLInitialResponse.Decode")
	*dst = SASLInitialResponse{}

	rp := 0

	idx := bytes.IndexByte(src, 0)
	if idx < 0 {
		return errors.New("invalid SASLInitialResponse")
	}

	dst.AuthMechanism = string(src[rp:idx])
	rp = idx + 1

	rp += 4 // The rest of the message is data so we can just skip the size
	dst.Data = src[rp:]

	return nil
}

// Encode encodes src into dst. dst will include the 1 byte message type identifier and the 4 byte message length.
func (src *SASLInitialResponse) Encode(dst []byte) []byte {
	//println("SASLInitialResponse.Encode")
	dst = append(dst, 'p')
	sp := len(dst)
	dst = pgio.AppendInt32(dst, -1)

	dst = append(dst, []byte(src.AuthMechanism)...)
	dst = append(dst, 0)

	dst = pgio.AppendInt32(dst, int32(len(src.Data)))
	dst = append(dst, src.Data...)

	pgio.SetInt32(dst[sp:], int32(len(dst[sp:])))

	return dst
}

// MarshalJSON implements encoding/json.Marshaler.
// func (src SASLInitialResponse) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Type          string
// 		AuthMechanism string
// 		Data          string
// 	}{
// 		Type:          "SASLInitialResponse",
// 		AuthMechanism: src.AuthMechanism,
// 		Data:          string(src.Data),
// 	})
// }

// UnmarshalJSON implements encoding/json.Unmarshaler.
// func (dst *SASLInitialResponse) UnmarshalJSON(data []byte) error {
// 	// Ignore null, like in the main JSON package.
// 	if string(data) == "null" {
// 		return nil
// 	}

// 	var msg struct {
// 		AuthMechanism string
// 		Data          string
// 	}
// 	if err := json.Unmarshal(data, &msg); err != nil {
// 		return err
// 	}
// 	dst.AuthMechanism = msg.AuthMechanism
// 	dst.Data = []byte(msg.Data)
// 	return nil
// }
