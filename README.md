ersatzjson
==========

Ersatzjson isn't really json, until it is. 

Json doesn't support comments since machines don't really have a use for them and why transmit the extra bytes?  However, humans like comments, for good reason.  Ersatzjson will read json with comments and return the json with the comments elided.

Ersatzjson supports two forms of comments: single-line, `\\`, and multi-line, `\* *\`. These work as you would expect comments to work.  For single-line comments, everything after a `\\` is ignored until a newline character is encountered, `\r`, `\n`, or `\r\n`, depending on your system. For multi-line comments, everything after a `\*` is ignored until a corresponding `(\` is encountered. 

Everything within single quotes, `' '`, and double quotes, `" "` are ignored.

## Import

    import ej "github.com/mohae/ersatzjson"

The `ej` is optional, I use it because I'm lazy.

