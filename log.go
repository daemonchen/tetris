package tetris

import "flag"

// TODO:
type logType struct {
}

var (
	logLevel = flag.Int64("level", 0, "log level")
)

func init() {
	flag.Parse()
}

func trace(fmt string, args ...interface{})    {}
func debug(fmt string, args ...interface{})    {}
func info(fmt string, args ...interface{})     {}
func warning(fmt string, args ...interface{})  {}
func critical(fmt string, args ...interface{}) {}
func fatal(fmt string, args ...interface{})    {}
