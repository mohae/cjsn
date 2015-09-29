package ersatzjson

import (
	"encoding/json"

	"github.com/mohae/nocomment"
)

type Decoder struct {
	json.Decoder
}

// Unmarshal elides comments from the "json" and then calls
// json.Unmarshal to unmarshal to the provided interface{}.
func Unmarshal(data []byte, v interface{}) error {
	data = nocomment.Clean(data)
	err := json.Unmarshal(data, &v)
	return err
}
