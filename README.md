CJSON
=====
[![Build Status](https://travis-ci.org/mohae/cjson.png)](https://travis-ci.org/mohae/cjson)

CJSON: commented JSON

JSON doesn't support comments since machines don't really have a use for them and why transmit the extra bytes?  However, humans like comments, for good reason.  CJSON unmarshals commented JSON.

By convention, I use either `cjson` or `cjsn` as the file extension.  The actual file extension doesn't matter as this package takes `[]byte`.  The api is consistent with `encoding/json`'s Unmarshal.

## Comments
[Nocomment](https://github.com/mohae/nocomment) is used to elide comments from the JSON. Once the comments have been elided, the resulting bytes are consistent with JSON as defined in RFC 4627 and are unmarshaled into the passed `interface{}` by `encoding/json`.

CJSON supports both block comments and line comments.

Two types of line comments are supported. Line comments can begin anywhere on a line and start with either `//` or `#`. Line comments end when an EOL is encountered, both `\n` and `\r\n` are interpreted as a newline.

Block comments start with `\*` and continue until a `*\` is encountered.

Comment sequences are not matched if they occur within double-quotes, `" "`.

## Usage
CJSON's `Unmarshal` has the same signature as `json.Unmarshal`; it wraps `encoding/json`'s Unmarshal.  When unmarshaling, the JSON first has it's comments elided.  Once cleaned, it is unmarshaled via `encoding/json`'s `Unmarshal`.

Import:

    import "github.com/mohae/cjson"

Unmarshal:

    var data MyStruct
    err := cjson.Unmarshal(jsonBlob, &data)
    if err != nil {
        log.Error(err)
        return
    }

## Notes
I use `cjsn` or `cjson` as the extension for JSON with comments.  While the actual extension for commented JSON files doesn't matter, it is recommended that it is not one of the common JSON extensions. Commented JSON isn't real JSON and using a JSON extension may cause problems or confusion.
