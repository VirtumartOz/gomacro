// this file was generated by gomacro command: import "flag"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"flag"
)

func init() {
	Binds["flag"] = map[string]Value{
		"Arg":	ValueOf(flag.Arg),
		"Args":	ValueOf(flag.Args),
		"Bool":	ValueOf(flag.Bool),
		"BoolVar":	ValueOf(flag.BoolVar),
		"CommandLine":	ValueOf(&flag.CommandLine).Elem(),
		"ContinueOnError":	ValueOf(flag.ContinueOnError),
		"Duration":	ValueOf(flag.Duration),
		"DurationVar":	ValueOf(flag.DurationVar),
		"ErrHelp":	ValueOf(&flag.ErrHelp).Elem(),
		"ExitOnError":	ValueOf(flag.ExitOnError),
		"Float64":	ValueOf(flag.Float64),
		"Float64Var":	ValueOf(flag.Float64Var),
		"Int":	ValueOf(flag.Int),
		"Int64":	ValueOf(flag.Int64),
		"Int64Var":	ValueOf(flag.Int64Var),
		"IntVar":	ValueOf(flag.IntVar),
		"Lookup":	ValueOf(flag.Lookup),
		"NArg":	ValueOf(flag.NArg),
		"NFlag":	ValueOf(flag.NFlag),
		"NewFlagSet":	ValueOf(flag.NewFlagSet),
		"PanicOnError":	ValueOf(flag.PanicOnError),
		"Parse":	ValueOf(flag.Parse),
		"Parsed":	ValueOf(flag.Parsed),
		"PrintDefaults":	ValueOf(flag.PrintDefaults),
		"Set":	ValueOf(flag.Set),
		"String":	ValueOf(flag.String),
		"StringVar":	ValueOf(flag.StringVar),
		"Uint":	ValueOf(flag.Uint),
		"Uint64":	ValueOf(flag.Uint64),
		"Uint64Var":	ValueOf(flag.Uint64Var),
		"UintVar":	ValueOf(flag.UintVar),
		"UnquoteUsage":	ValueOf(flag.UnquoteUsage),
		"Usage":	ValueOf(&flag.Usage).Elem(),
		"Var":	ValueOf(flag.Var),
		"Visit":	ValueOf(flag.Visit),
		"VisitAll":	ValueOf(flag.VisitAll),
	}
	Types["flag"] = map[string]Type{
		"ErrorHandling":	TypeOf((*flag.ErrorHandling)(nil)).Elem(),
		"Flag":	TypeOf((*flag.Flag)(nil)).Elem(),
		"FlagSet":	TypeOf((*flag.FlagSet)(nil)).Elem(),
		"Getter":	TypeOf((*flag.Getter)(nil)).Elem(),
		"Value":	TypeOf((*flag.Value)(nil)).Elem(),
	}
	Proxies["flag"] = map[string]Type{
		"Getter":	TypeOf((*Getter_flag)(nil)).Elem(),
		"Value":	TypeOf((*Value_flag)(nil)).Elem(),
	}
}

// --------------- proxy for flag.Getter ---------------
type Getter_flag struct {
	Get_	func() interface{}
	Set_	func(string) error
	String_	func() string
}
func (Obj Getter_flag) Get() interface{} {
	return Obj.Get_()
}
func (Obj Getter_flag) Set(unnamed0 string) error {
	return Obj.Set_(unnamed0)
}
func (Obj Getter_flag) String() string {
	return Obj.String_()
}

// --------------- proxy for flag.Value ---------------
type Value_flag struct {
	Set_	func(string) error
	String_	func() string
}
func (Obj Value_flag) Set(unnamed0 string) error {
	return Obj.Set_(unnamed0)
}
func (Obj Value_flag) String() string {
	return Obj.String_()
}
