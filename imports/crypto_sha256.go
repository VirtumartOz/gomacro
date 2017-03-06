// this file was generated by gomacro command: import "crypto/sha256"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/sha256"
)

func init() {
	Binds["crypto/sha256"] = map[string]Value{
		"BlockSize":	ValueOf(sha256.BlockSize),
		"New":	ValueOf(sha256.New),
		"New224":	ValueOf(sha256.New224),
		"Size":	ValueOf(sha256.Size),
		"Size224":	ValueOf(sha256.Size224),
		"Sum224":	ValueOf(sha256.Sum224),
		"Sum256":	ValueOf(sha256.Sum256),
	}
	Types["crypto/sha256"] = map[string]Type{
	}
	Proxies["crypto/sha256"] = map[string]Type{
	}
}
