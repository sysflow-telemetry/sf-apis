//
// Copyright (C) 2021 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logger implements logging utilities.
package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// LogLevel type for enumeration.
type LogLevel int

// LogLevel enumeration.
const (
	TRACE LogLevel = iota
	INFO
	WARN
	ERROR
	HEALTH
	QUIET
)

func (d LogLevel) String() string {
	return [...]string{"Trace", "Info", "Warn", "Error", "Health", "Quiet"}[d]
}

// GetLogLevelFromValue returns LogLevel corresponding to string s (if not found, defaults to INFO).
func GetLogLevelFromValue(s string) LogLevel {
	m := map[string]LogLevel{"trace": TRACE, "info": INFO, "warn": WARN, "error": ERROR, "health": HEALTH, "quiet": QUIET}
	if l, ok := m[strings.ToLower(s)]; ok {
		return l
	}
	return INFO
}

// Loggers reflecting different log levels.
var (
	Trace  *log.Logger
	Info   *log.Logger
	Warn   *log.Logger
	Error  *log.Logger
	Health *log.Logger
)

// InitLoggers initialize utility loggers with default i/o streams.
func InitLoggers(level LogLevel) {
	switch level {
	case TRACE:
		initLoggers(os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stdout)
	case INFO:
		initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, os.Stdout)
	case WARN:
		initLoggers(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr, os.Stdout)
	case ERROR:
		initLoggers(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr, os.Stdout)
	case HEALTH:
		initLoggers(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stdout)
	case QUIET:
		initLoggers(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
	default:
		initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, os.Stdout)
	}
}

func initLoggers(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warnHandle io.Writer,
	errorHandle io.Writer,
	healthHandle io.Writer) {

	Trace = log.New(traceHandle,
		fmt.Sprintf("[%s] ", TRACE),
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		fmt.Sprintf("[%s] ", INFO),
		log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(warnHandle,
		fmt.Sprintf("[%s] ", WARN),
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		fmt.Sprintf("[%s] ", ERROR),
		log.Ldate|log.Ltime|log.Lshortfile)

	Health = log.New(healthHandle,
		fmt.Sprintf("[%s] ", HEALTH),
		log.Ldate|log.Ltime|log.Lshortfile)
}
