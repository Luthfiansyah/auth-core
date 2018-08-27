package logging

import (
	"os"

	opLogging "github.com/op/go-logging"
)

var log = opLogging.MustGetLogger("auth")

const INTERNAL = "internal"

func MustGetLogger(name string) *AuthLogger {
	host, err := os.Hostname()
	if err != nil {
		log.Error("", err.Error())
		host = "unknown"
	}
	ppl := &AuthLogger{opLogging.MustGetLogger(name), host}
	ppl.ExtraCalldepth = 1
	return ppl
}

type AuthLogger struct {
	*opLogging.Logger
	Hostname string
}

// func (ppl *AuthLogger) Debug(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Debug(args...)
// }
//
// func (ppl *AuthLogger) Debugf(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Debugf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Info(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Info(args)
// }
//
// func (ppl *AuthLogger) Infof(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Infof("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Error(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Error(args)
// }
//
// func (ppl *AuthLogger) Errorf(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Errorf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Critical(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Critical(args)
// }
//
// func (ppl *AuthLogger) Criticalf(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Criticalf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Fatal(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Fatal(args)
// }
//
// func (ppl *AuthLogger) Fatalf(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Fatalf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Panic(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Panic(args)
// }
//
// func (ppl *AuthLogger) Panicf(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Panicf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Warning(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Warning(args)
// }
//
// func (ppl *AuthLogger) Warningf(userid string, string_format string, args ...interface{}) {
// 	ppl.Logger.Warningf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
// }
//
// func (ppl *AuthLogger) Notice(userid string, args ...interface{}) {
// 	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
// 	ppl.Logger.Notice(args)
// }

func (ppl *AuthLogger) Noticef(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Noticef("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}
