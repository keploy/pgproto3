package pgproto3

import (
	"encoding/binary"

	"github.com/jackc/pgio"
)

type BackendKeyData struct {
	ProcessID uint32 `json:"process_id" yaml:"process_id"`
	SecretKey uint32 `json:"secret_key" yaml:"secret_key"`
}

// Backend identifies this message as sendable by the PostgreSQL backend.
func (*BackendKeyData) Backend() {}

// Decode decodes src into dst. src must contain the complete message with the exception of the initial 1 byte message
// type identifier and 4 byte message length.
func (dst *BackendKeyData) Decode(src []byte) error {
	//println("BackendKeyData.Decode")
	if len(src) != 8 {
		return &invalidMessageLenErr{messageType: "BackendKeyData", expectedLen: 8, actualLen: len(src)}
	}

	dst.ProcessID = binary.BigEndian.Uint32(src[:4])
	dst.SecretKey = binary.BigEndian.Uint32(src[4:])

	return nil
}

// Encode encodes src into dst. dst will include the 1 byte message type identifier and the 4 byte message length.
func (src *BackendKeyData) Encode(dst []byte) []byte {
	//println("BackendKeyData.Encode")
	dst = append(dst, 'K')
	dst = pgio.AppendUint32(dst, 12)
	dst = pgio.AppendUint32(dst, src.ProcessID)
	dst = pgio.AppendUint32(dst, src.SecretKey)
	return dst
}

// // MarshalJSON implements encoding/json.Marshaler.
// func (src BackendKeyData) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Type      string
// 		ProcessID uint32
// 		SecretKey uint32
// 	}{
// 		Type:      "BackendKeyData",
// 		ProcessID: src.ProcessID,
// 		SecretKey: src.SecretKey,
// 	})
// }
