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

package cjson

import (
	"testing"
)

type Animal struct {
	Name  string
	Order string
}

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		jsonBlob    []byte
		expected    []Animal
		expectedErr string
	}{
		{jsonBlob: []byte(""), expected: []Animal{}, expectedErr: "unexpected end of JSON input"},
		// from http://golang.org/pkg/encoding/json/#example_Unmarshal
		{jsonBlob: []byte(`[
      {"Name": "Platypus", "Order": "Monotremata"},
      {"Name": "Quoll", "Order": "Dasyuromorphia"}
    ]`), expected: []Animal{
			Animal{"Platypus", "Monotremata"},
			Animal{"Quoll", "Dasyuromorphia"},
		}, expectedErr: ""},
		{jsonBlob: []byte(`// comments
      [
      {"Name": "Platypus", "Order": "Monotremata"},
      {"Name": "Quoll", "Order": "Dasyuromorphia"}
    ]`), expected: []Animal{
			Animal{"Platypus", "Monotremata"},
			Animal{"Quoll", "Dasyuromorphia"},
		}, expectedErr: ""},
		{jsonBlob: []byte(`[ /* block comment */
      {"Name": "Platypus", "Order": "Monotremata"},
      # another one
      {"Name": "Quoll", "Order": "Dasyuromorphia"}
    ]`), expected: []Animal{
			Animal{"Platypus", "Monotremata"},
			Animal{"Quoll", "Dasyuromorphia"},
		}, expectedErr: ""},
	}

	for i, test := range tests {
		var animals []Animal
		err := Unmarshal(test.jsonBlob, &animals)
		if err != nil {
			if err.Error() != test.expectedErr {
				t.Errorf("%d: expected err %q got %q", i, test.expectedErr, err)
			}
			continue
		}
		if len(test.expected) != len(animals) {
			t.Errorf("%d: expected %d results, go %d", i, len(test.expected), len(animals))
			continue
		}
		for i, v := range animals {
			j, found := findAnimal(v, test.expected)
			if !found {
				t.Errorf("%d: unexpected result: %q", i, v.Name)
				continue
			}
			if test.expected[j].Order != v.Order {
				t.Errorf("%d expected the order for %q to be %q, got %q", i, v.Name, test.expected[j].Order, v.Order)
			}
		}
	}
}

func findAnimal(val Animal, animals []Animal) (int, bool) {
	for i, v := range animals {
		if val.Name == v.Name {
			return i, true
		}
	}
	return 0, false
}
