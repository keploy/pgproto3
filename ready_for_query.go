package pgproto3

type ReadyForQuery struct {
	TxStatus byte `json:"tx_status"`
}

// Backend identifies this message as sendable by the PostgreSQL backend.
func (*ReadyForQuery) Backend() {}

// Decode decodes src into dst. src must contain the complete message with the exception of the initial 1 byte message
// type identifier and 4 byte message length.
func (dst *ReadyForQuery) Decode(src []byte) error {
	//println("ReadyForQuery.Decode")
	if len(src) != 1 {
		return &invalidMessageLenErr{messageType: "ReadyForQuery", expectedLen: 1, actualLen: len(src)}
	}

	dst.TxStatus = src[0]

	return nil
}

// Encode encodes src into dst. dst will include the 1 byte message type identifier and the 4 byte message length.
func (src *ReadyForQuery) Encode(dst []byte) []byte {
	//println("ReadyForQuery.Encode")
	return append(dst, 'Z', 0, 0, 0, 5, src.TxStatus)
}

// // MarshalJSON implements encoding/json.Marshaler.
// func (src ReadyForQuery) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Type     string
// 		TxStatus string
// 	}{
// 		Type:     "ReadyForQuery",
// 		TxStatus: string(src.TxStatus),
// 	})
// }

// // UnmarshalJSON implements encoding/json.Unmarshaler.
// func (dst *ReadyForQuery) UnmarshalJSON(data []byte) error {
// 	// Ignore null, like in the main JSON package.
// 	if string(data) == "null" {
// 		return nil
// 	}

// 	var msg struct {
// 		TxStatus string
// 	}
// 	if err := json.Unmarshal(data, &msg); err != nil {
// 		return err
// 	}
// 	if len(msg.TxStatus) != 1 {
// 		return errors.New("invalid length for ReadyForQuery.TxStatus")
// 	}
// 	dst.TxStatus = msg.TxStatus[0]
// 	return nil
// }
