// this file was generated by gomacro command: import "compress/gzip"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"compress/gzip"
)

func init() {
	Binds["compress/gzip"] = map[string]Value{
		"BestCompression":	ValueOf(gzip.BestCompression),
		"BestSpeed":	ValueOf(gzip.BestSpeed),
		"DefaultCompression":	ValueOf(gzip.DefaultCompression),
		"ErrChecksum":	ValueOf(&gzip.ErrChecksum).Elem(),
		"ErrHeader":	ValueOf(&gzip.ErrHeader).Elem(),
		"HuffmanOnly":	ValueOf(gzip.HuffmanOnly),
		"NewReader":	ValueOf(gzip.NewReader),
		"NewWriter":	ValueOf(gzip.NewWriter),
		"NewWriterLevel":	ValueOf(gzip.NewWriterLevel),
		"NoCompression":	ValueOf(gzip.NoCompression),
	}
	Types["compress/gzip"] = map[string]Type{
		"Header":	TypeOf((*gzip.Header)(nil)).Elem(),
		"Reader":	TypeOf((*gzip.Reader)(nil)).Elem(),
		"Writer":	TypeOf((*gzip.Writer)(nil)).Elem(),
	}
	Proxies["compress/gzip"] = map[string]Type{
	}
}
