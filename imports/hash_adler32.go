// this file was generated by gomacro command: import "hash/adler32"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"hash/adler32"
)

func init() {
	Binds["hash/adler32"] = map[string]Value{
		"Checksum":	ValueOf(adler32.Checksum),
		"New":	ValueOf(adler32.New),
		"Size":	ValueOf(adler32.Size),
	}
	Types["hash/adler32"] = map[string]Type{
	}
	Proxies["hash/adler32"] = map[string]Type{
	}
}
