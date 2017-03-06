// this file was generated by gomacro command: import "fmt"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"fmt"
)

func init() {
	Binds["fmt"] = map[string]Value{
		"Errorf":	ValueOf(fmt.Errorf),
		"Fprint":	ValueOf(fmt.Fprint),
		"Fprintf":	ValueOf(fmt.Fprintf),
		"Fprintln":	ValueOf(fmt.Fprintln),
		"Fscan":	ValueOf(fmt.Fscan),
		"Fscanf":	ValueOf(fmt.Fscanf),
		"Fscanln":	ValueOf(fmt.Fscanln),
		"Print":	ValueOf(fmt.Print),
		"Printf":	ValueOf(fmt.Printf),
		"Println":	ValueOf(fmt.Println),
		"Scan":	ValueOf(fmt.Scan),
		"Scanf":	ValueOf(fmt.Scanf),
		"Scanln":	ValueOf(fmt.Scanln),
		"Sprint":	ValueOf(fmt.Sprint),
		"Sprintf":	ValueOf(fmt.Sprintf),
		"Sprintln":	ValueOf(fmt.Sprintln),
		"Sscan":	ValueOf(fmt.Sscan),
		"Sscanf":	ValueOf(fmt.Sscanf),
		"Sscanln":	ValueOf(fmt.Sscanln),
	}
	Types["fmt"] = map[string]Type{
		"Formatter":	TypeOf((*fmt.Formatter)(nil)).Elem(),
		"GoStringer":	TypeOf((*fmt.GoStringer)(nil)).Elem(),
		"ScanState":	TypeOf((*fmt.ScanState)(nil)).Elem(),
		"Scanner":	TypeOf((*fmt.Scanner)(nil)).Elem(),
		"State":	TypeOf((*fmt.State)(nil)).Elem(),
		"Stringer":	TypeOf((*fmt.Stringer)(nil)).Elem(),
	}
	Proxies["fmt"] = map[string]Type{
		"Formatter":	TypeOf((*Formatter_fmt)(nil)).Elem(),
		"GoStringer":	TypeOf((*GoStringer_fmt)(nil)).Elem(),
		"ScanState":	TypeOf((*ScanState_fmt)(nil)).Elem(),
		"Scanner":	TypeOf((*Scanner_fmt)(nil)).Elem(),
		"State":	TypeOf((*State_fmt)(nil)).Elem(),
		"Stringer":	TypeOf((*Stringer_fmt)(nil)).Elem(),
	}
}

// --------------- proxy for fmt.Formatter ---------------
type Formatter_fmt struct {
	Format_	func(f fmt.State, c rune) 
}
func (Obj Formatter_fmt) Format(f fmt.State, c rune)  {
	Obj.Format_(f, c)
}

// --------------- proxy for fmt.GoStringer ---------------
type GoStringer_fmt struct {
	GoString_	func() string
}
func (Obj GoStringer_fmt) GoString() string {
	return Obj.GoString_()
}

// --------------- proxy for fmt.ScanState ---------------
type ScanState_fmt struct {
	Read_	func(buf []byte) (n int, err error)
	ReadRune_	func() (r rune, size int, err error)
	SkipSpace_	func() 
	Token_	func(skipSpace bool, f func(rune) bool) (token []byte, err error)
	UnreadRune_	func() error
	Width_	func() (wid int, ok bool)
}
func (Obj ScanState_fmt) Read(buf []byte) (n int, err error) {
	return Obj.Read_(buf)
}
func (Obj ScanState_fmt) ReadRune() (r rune, size int, err error) {
	return Obj.ReadRune_()
}
func (Obj ScanState_fmt) SkipSpace()  {
	Obj.SkipSpace_()
}
func (Obj ScanState_fmt) Token(skipSpace bool, f func(rune) bool) (token []byte, err error) {
	return Obj.Token_(skipSpace, f)
}
func (Obj ScanState_fmt) UnreadRune() error {
	return Obj.UnreadRune_()
}
func (Obj ScanState_fmt) Width() (wid int, ok bool) {
	return Obj.Width_()
}

// --------------- proxy for fmt.Scanner ---------------
type Scanner_fmt struct {
	Scan_	func(state fmt.ScanState, verb rune) error
}
func (Obj Scanner_fmt) Scan(state fmt.ScanState, verb rune) error {
	return Obj.Scan_(state, verb)
}

// --------------- proxy for fmt.State ---------------
type State_fmt struct {
	Flag_	func(c int) bool
	Precision_	func() (prec int, ok bool)
	Width_	func() (wid int, ok bool)
	Write_	func(b []byte) (n int, err error)
}
func (Obj State_fmt) Flag(c int) bool {
	return Obj.Flag_(c)
}
func (Obj State_fmt) Precision() (prec int, ok bool) {
	return Obj.Precision_()
}
func (Obj State_fmt) Width() (wid int, ok bool) {
	return Obj.Width_()
}
func (Obj State_fmt) Write(b []byte) (n int, err error) {
	return Obj.Write_(b)
}

// --------------- proxy for fmt.Stringer ---------------
type Stringer_fmt struct {
	String_	func() string
}
func (Obj Stringer_fmt) String() string {
	return Obj.String_()
}
