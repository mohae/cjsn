// Copyright 2017 Joel Scoble
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

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
