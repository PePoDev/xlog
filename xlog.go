// Package xlog is wrapper function for logrus api
package xlog

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logFormat *logrus.TextFormatter

func init() {
	DefaultLogFormatter()
}

// DefaultLogFormatter set Formatter with TextFormatter that log like.
// You can use logrus package to set it with yourself
func DefaultLogFormatter() {
	logFormat = &logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		DisableColors:             false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "01/02/2006 03:04:05 PM",
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          true,
		FieldMap: logrus.FieldMap{
			"time":  "Timestamp",
			"level": "Level",
			"msg":   "Message",
		},
	}

	logrus.SetFormatter(logFormat)
	logrus.SetReportCaller(false)

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

// GetFormatter return Formatter that use by log for do mini custom by yourself
func GetFormatter() logrus.Formatter {
	return logFormat
}

// SetFormatter custom logrus Formatter
func SetFormatter(format logrus.Formatter) {
	logrus.SetFormatter(format)
}

// SetLevel will set logrus level
func SetLevel(level logrus.Level) {
	logrus.SetLevel(level)

}

// GetLevel returns the standard logger level.
func GetLevel() logrus.Level {
	return logrus.GetLevel()
}

// Log struct contain logrus fields
type Log struct {
	fields logrus.Fields
}

// NewLog Return new log struct
func NewLog() *Log {
	log := &Log{}
	log.fields = make(map[string]interface{})
	return log
}

func (log Log) attachTrace() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	log.fields["Caller"] = fmt.Sprintf("%v:%v [%v]", frame.File, frame.Line, frame.Function)
}

// SetField for set log tag that you want
func (log Log) SetField(fieldName string, value interface{}) *Log {
	log.fields[fieldName] = value
	return &log
}

// Tracef will log in Trace level with format
func (log Log) Tracef(format string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Tracef(format, args...)
}

// Traceln will log in Trace level with format
func (log Log) Traceln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Traceln(args...)
}

// Trace will log in Trace level
func (log Log) Trace(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Trace(message)
}

// Debugf will log in Debug level with format
func (log Log) Debugf(format string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Debugf(format, args...)
}

// Debugln will log in Debug level with format
func (log Log) Debugln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Debugln(args...)
}

// Debug will log in Debug level
func (log Log) Debug(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Debug(message)
}

// Infof will log in info level with format
func (log Log) Infof(format string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Infof(format, args...)
}

// Infoln will log in info level with format
func (log Log) Infoln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Infoln(args...)
}

// Info will log in info level
func (log Log) Info(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Info(message)
}

// Warnf will log in Warning level with format
func (log Log) Warnf(format string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Warnf(format, args...)
}

// Warnln will log in Warning level with format
func (log Log) Warnln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Warnln(args...)
}

// Warn will log in Warning level
func (log Log) Warn(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Warn(message)
}

// Errorf will log in Error level with format
func (log Log) Errorf(format string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Errorf(format, args...)
}

// Errorln will log in Error level with format
func (log Log) Errorln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Errorln(args...)
}

// Error will log in Error level
func (log Log) Error(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Error(message)
}

// Fatalf will log in Fail level with format and call exit program
func (log Log) Fatalf(format string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Fatalf(format, args...)
}

// Fatalln will log in Fail level with format and call exit program
func (log Log) Fatalln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Fatalln(args...)
}

// Fatal will log in Fail level and call exit program
func (log Log) Fatal(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Fatal(message)
}

// Panicf will log in panic level with format and call panic
func (log Log) Panicf(message string, args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Panicf(message, args...)
}

// Panicln will log in panic level with format and call panic
func (log Log) Panicln(args ...interface{}) {
	log.attachTrace()
	logrus.WithFields(log.fields).Panicln(args...)
}

// Panic will log in panic level and call panic
func (log Log) Panic(message string) {
	log.attachTrace()
	logrus.WithFields(log.fields).Panic(message)
}

// Tracef will log in Trace level with format
func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

// Traceln will log in Trace level with format
func Traceln(args ...interface{}) {
	logrus.Traceln(args...)
}

// Trace will log in Trace level
func Trace(message string) {
	logrus.Trace(message)
}

// Debugf will log in Debug level with format
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Debugln will log in Debug level with format
func Debugln(args ...interface{}) {
	logrus.Debugln(args...)
}

// Debug will log in Debug level
func Debug(message string) {
	logrus.Debug(message)
}

// Infof will log in info level with format
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Infoln will log in info level with format
func Infoln(args ...interface{}) {
	logrus.Infoln(args...)
}

// Info will log in info level
func Info(message string) {
	logrus.Info(message)
}

// Warnf will log in Warning level with format
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Warnln will log in Warning level with format
func Warnln(args ...interface{}) {
	logrus.Warnln(args...)
}

// Warn will log in Warning level
func Warn(message string) {
	logrus.Warn(message)
}

// Errorf will log in Error level with format
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Errorln will log in Error level with format
func Errorln(args ...interface{}) {
	logrus.Errorln(args...)
}

// Error will log in Error level
func Error(message string) {
	logrus.Error(message)
}

// Fatalf will log in Fail level with format and call exit program
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Fatalln will log in Fail level with format and call exit program
func Fatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}

// Fatal will log in Fail level and call exit program
func Fatal(message string) {
	logrus.Fatal(message)
}

// Panicf will log in panic level with format and call panic
func Panicf(message string, args ...interface{}) {
	logrus.Panicf(message, args...)
}

// Panicln will log in panic level with format and call panic
func Panicln(args ...interface{}) {
	logrus.Panicln(args...)
}

// Panic will log in panic level and call panic
func Panic(message string) {
	logrus.Panic(message)
}
