cjsn
==========
[![Build Status](https://travis-ci.org/mohae/cjsn.png)](https://travis-ci.org/mohae/cjsn)

cjsn: commented JSON

JSON doesn't support comments since machines don't really have a use for them and why transmit the extra bytes?  However, humans like comments, for good reason.  cjsn unmarshals commented JSON. 

For now, that is all cjsn does. 

## Comments
[Nocomment](https://github.com/mohae/nocomment) is used to elide comments from the JSON. Once the comments have been elided, the resulting bytes are consistent with JSON as defined in RFC 4627 and are unmarshaled into the passed `interface{}` by `encoding/json`. 

cjsn supports both block comments and line comments.

Two types of line comments are supported. Line comments can begin anywhere on a line and start with either `//` or `#`. Line comments end when an EOL is encountered: `\r`, `\n`, or `\r\n` are all interpreted as a newline.

Block comments start with `\*` and continue until a `*\` is encountered. 

Comment sequences are not matched if they occur within double-quotes, `" "`.

## Usage
cjsn's `Unmarshal` has the same signature as `json.Unmarshal`; it wraps `encoding/json`'s Unmarshal.  When unmarshaling, the JSON first has it's comments elided.  Once cleaned, it is unmarshaled via `encoding/json`'s `Unmarshal`.

Import:

    import "github.com/mohae/cjsn"

Unmarshal:

    var data MyStruct
    err := cjsn.Unmarshal(jsonBlob, &data)
    if err != nil {
        log.Error(err)
        return
    }

## Notes
I use `cjsn` as the extension for JSON with comments.  While the actual extension for commented JSON files doesn't matter, it is recommended that it is not one of the common JSON extensions. Commented JSON isn't real JSON and using a JSON extension may cause problems or confusion.

## Wishlist

* Unmarshal json, preserving map order.  
* Marshal json, preserving map order.  

