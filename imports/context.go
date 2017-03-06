// this file was generated by gomacro command: import "context"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"context"
	"time"
)

func init() {
	Binds["context"] = map[string]Value{
		"Background":	ValueOf(context.Background),
		"Canceled":	ValueOf(&context.Canceled).Elem(),
		"DeadlineExceeded":	ValueOf(&context.DeadlineExceeded).Elem(),
		"TODO":	ValueOf(context.TODO),
		"WithCancel":	ValueOf(context.WithCancel),
		"WithDeadline":	ValueOf(context.WithDeadline),
		"WithTimeout":	ValueOf(context.WithTimeout),
		"WithValue":	ValueOf(context.WithValue),
	}
	Types["context"] = map[string]Type{
		"CancelFunc":	TypeOf((*context.CancelFunc)(nil)).Elem(),
		"Context":	TypeOf((*context.Context)(nil)).Elem(),
	}
	Proxies["context"] = map[string]Type{
		"Context":	TypeOf((*Context_context)(nil)).Elem(),
	}
}

// --------------- proxy for context.Context ---------------
type Context_context struct {
	Deadline_	func() (deadline time.Time, ok bool)
	Done_	func() <-chan struct{}
	Err_	func() error
	Value_	func(key interface{}) interface{}
}
func (Obj Context_context) Deadline() (deadline time.Time, ok bool) {
	return Obj.Deadline_()
}
func (Obj Context_context) Done() <-chan struct{} {
	return Obj.Done_()
}
func (Obj Context_context) Err() error {
	return Obj.Err_()
}
func (Obj Context_context) Value(key interface{}) interface{} {
	return Obj.Value_(key)
}
