// this file was generated by gomacro command: import "go/importer"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"go/importer"
)

func init() {
	Binds["go/importer"] = map[string]Value{
		"Default":	ValueOf(importer.Default),
		"For":	ValueOf(importer.For),
	}
	Types["go/importer"] = map[string]Type{
		"Lookup":	TypeOf((*importer.Lookup)(nil)).Elem(),
	}
	Proxies["go/importer"] = map[string]Type{
	}
}
