// this file was generated by gomacro command: import "crypto/hmac"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"crypto/hmac"
)

func init() {
	Binds["crypto/hmac"] = map[string]Value{
		"Equal":	ValueOf(hmac.Equal),
		"New":	ValueOf(hmac.New),
	}
	Types["crypto/hmac"] = map[string]Type{
	}
	Proxies["crypto/hmac"] = map[string]Type{
	}
}
