package pgproto3

import (
	"github.com/jackc/pgio"
)

type GSSResponse struct {
	Data []byte `json:"data" yaml:"data"`
}

// Frontend identifies this message as sendable by a PostgreSQL frontend.
func (g *GSSResponse) Frontend() {}

func (g *GSSResponse) Decode(data []byte) error {
	//println("GSSResponse.Decode")
	g.Data = data
	return nil
}

func (g *GSSResponse) Encode(dst []byte) []byte {
	//println("GSSResponse.Encode")
	dst = append(dst, 'p')
	dst = pgio.AppendInt32(dst, int32(4+len(g.Data)))
	dst = append(dst, g.Data...)
	return dst
}

// // MarshalJSON implements encoding/json.Marshaler.
// func (g *GSSResponse) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Type string
// 		Data []byte
// 	}{
// 		Type: "GSSResponse",
// 		Data: g.Data,
// 	})
// }

// // UnmarshalJSON implements encoding/json.Unmarshaler.
// func (g *GSSResponse) UnmarshalJSON(data []byte) error {
// 	var msg struct {
// 		Data []byte
// 	}
// 	if err := json.Unmarshal(data, &msg); err != nil {
// 		return err
// 	}
// 	g.Data = msg.Data
// 	return nil
// }
