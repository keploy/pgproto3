package pgproto3

type NoticeResponse ErrorResponse

// Backend identifies this message as sendable by the PostgreSQL backend.
func (*NoticeResponse) Backend() {}

// Decode decodes src into dst. src must contain the complete message with the exception of the initial 1 byte message
// type identifier and 4 byte message length.
func (dst *NoticeResponse) Decode(src []byte) error {
	//println("NoticeResponse.Decode")
	return (*ErrorResponse)(dst).Decode(src)
}

// Encode encodes src into dst. dst will include the 1 byte message type identifier and the 4 byte message length.
func (src *NoticeResponse) Encode(dst []byte) []byte {
	//println("NoticeResponse.Encode")
	return append(dst, (*ErrorResponse)(src).marshalBinary('N')...)
}
