// this file was generated by gomacro command: import "net/smtp"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"net/smtp"
)

func init() {
	Binds["net/smtp"] = map[string]Value{
		"CRAMMD5Auth":	ValueOf(smtp.CRAMMD5Auth),
		"Dial":	ValueOf(smtp.Dial),
		"NewClient":	ValueOf(smtp.NewClient),
		"PlainAuth":	ValueOf(smtp.PlainAuth),
		"SendMail":	ValueOf(smtp.SendMail),
	}
	Types["net/smtp"] = map[string]Type{
		"Auth":	TypeOf((*smtp.Auth)(nil)).Elem(),
		"Client":	TypeOf((*smtp.Client)(nil)).Elem(),
		"ServerInfo":	TypeOf((*smtp.ServerInfo)(nil)).Elem(),
	}
	Proxies["net/smtp"] = map[string]Type{
		"Auth":	TypeOf((*Auth_net_smtp)(nil)).Elem(),
	}
}

// --------------- proxy for net/smtp.Auth ---------------
type Auth_net_smtp struct {
	Next_	func(fromServer []byte, more bool) (toServer []byte, err error)
	Start_	func(server *smtp.ServerInfo) (proto string, toServer []byte, err error)
}
func (Obj Auth_net_smtp) Next(fromServer []byte, more bool) (toServer []byte, err error) {
	return Obj.Next_(fromServer, more)
}
func (Obj Auth_net_smtp) Start(server *smtp.ServerInfo) (proto string, toServer []byte, err error) {
	return Obj.Start_(server)
}
