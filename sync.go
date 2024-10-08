package pgproto3

type Sync struct{}

// Frontend identifies this message as sendable by a PostgreSQL frontend.
func (*Sync) Frontend() {}

// Decode decodes src into dst. src must contain the complete message with the exception of the initial 1 byte message
// type identifier and 4 byte message length.
func (dst *Sync) Decode(src []byte) error {
	//println("Sync.Decode")
	if len(src) != 0 {
		// dst.isSync = true
		return &invalidMessageLenErr{messageType: "Sync", expectedLen: 0, actualLen: len(src)}
	}
	// dst.isSync = false
	return nil
}

// Encode encodes src into dst. dst will include the 1 byte message type identifier and the 4 byte message length.
func (src *Sync) Encode(dst []byte) []byte {
	// //println("Sync.Encode")
	return append(dst, 'S', 0, 0, 0, 4)
}

// MarshalJSON implements encoding/json.Marshaler.
// func (src Sync) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Type string
// 	}{
// 		Type: "Sync",
// 	})
// }
