// this file was generated by gomacro command: import "net/http/fcgi"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/http/fcgi"
)

func init() {
	Binds["net/http/fcgi"] = map[string]Value{
		"ErrConnClosed":	ValueOf(&fcgi.ErrConnClosed).Elem(),
		"ErrRequestAborted":	ValueOf(&fcgi.ErrRequestAborted).Elem(),
		"Serve":	ValueOf(fcgi.Serve),
	}
	Types["net/http/fcgi"] = map[string]Type{
	}
	Proxies["net/http/fcgi"] = map[string]Type{
	}
}
