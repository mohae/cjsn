// Package cjson supports json with comments: commented json.
//
// Currently, it only implements an Unmarshal, which accepts json as []byte and
// elides any comments before unmarshaling.  Comments can be either line
// comments, # and //, or block comments, /* */.
package cjson

import (
	"encoding/json"

	"github.com/mohae/nocomment"
)

// Decoder is a wrapper to json.Decoder
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
