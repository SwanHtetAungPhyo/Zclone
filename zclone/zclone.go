package Zclone

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

type LogLevel string

const (
	LogDebug LogLevel = "DEBUG"
	LogInfo  LogLevel = "INFO"
	LogTrace LogLevel = "TRACE"
	LogError LogLevel = "ERROR"
)

type Field struct {
	Key   string
	Value any
}

func (f *Field) FieldMarshal() ([]byte, error) {
	return json.Marshal(map[string]any{
		f.Key: f.Value,
	})
}

type LogEntry struct {
	Level     LogLevel
	TimeStamp string
	Message   string
	Caller    string
	Fields    []Field
}

type Logger struct {
	output   []io.Writer
	logLevel LogLevel
}

// func(Log.INFO, os.stdin, fileName)
func NewLogger(loglevel LogLevel, output ...io.Writer) *Logger {
	newlogger := &Logger{
		logLevel: loglevel,
	}
	if len(output) == 0 {
		newlogger.output = append(newlogger.output, os.Stdout)
	}

	// if len(output) > 0 {
	// 	newlogger.output[0] = output[0]
	// 	newlogger.output[1] = output[1]
	// }
	return newlogger
}

func (l *Logger) log(level LogLevel, message string, fields ...Field) {
	if !l.shouldLog(level) {
		return
	}
	logEntry := &LogEntry{
		Level:     level,
		TimeStamp: time.Now().Format("2006-01-02T15:04:05ZO7:00"),
		Message:   message,
		Caller:    l.currentFunctionName(),
		Fields:    fields,
	}
	data, err := json.MarshalIndent(logEntry, " ", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in Marshalling the log entries...%s", err.Error())
		return
	}
	fmt.Fprintln(l.output[0], string(data))
}

func (l *Logger) Info(logLevel LogLevel, message string, fields ...Field) {
	l.log(logLevel, message, fields...)
}

func (l *Logger) Debug(logLevel LogLevel, message string, fields ...Field) {
	l.log(logLevel, message, fields...)
}

func (l *Logger) Error(logLevel LogLevel, message string, fields ...Field) {
	l.log(logLevel, message, fields...)
}

func (l *Logger) Trace(logLevel LogLevel, message string, fields ...Field) {
	l.log(logLevel, message, fields...)
}
func (l *Logger) currentFunctionName() string {
	pc, fileName, lineNumber, ok := runtime.Caller(2)
	if !ok {
		return "Unknown function stack"
	}
	funcName := runtime.FuncForPC(pc)
	formattedString := fmt.Sprintf("fileName: %s | lineNumber: %d |functionName: %s", fileName, lineNumber, funcName.Name())
	return formattedString

}
func (l *Logger) shouldLog(level LogLevel) bool {
	logMap := map[LogLevel]int{
		LogDebug: 1,
		LogInfo:  2,
		LogTrace: 3,
		LogError: 4,
	}
	return logMap[level] >= logMap[l.logLevel]
}

// func (l *Logger) Close(){
// 	if l.output
// }

func String(key string, value any) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

func Int(key string, value any) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}

func Bool(key string, value any) Field {
	return Field{
		Key:   key,
		Value: value,
	}
}
