package log15

import (
	"os"

	"github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"
)

// Predefined handlers
var (
	root          *logger
	StdoutHandler = StreamHandler(os.Stdout, LogfmtFormat())
	StderrHandler = StreamHandler(os.Stderr, LogfmtFormat())
)

func init() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		StdoutHandler = StreamHandler(colorable.NewColorableStdout(), TerminalFormat())
	}

	if isatty.IsTerminal(os.Stderr.Fd()) {
		StderrHandler = StreamHandler(colorable.NewColorableStderr(), TerminalFormat())
	}

	root = &logger{[]interface{}{}, new(swapHandler)}
	root.SetHandler(StdoutHandler)
}

// New returns a new logger with the given context.
// New is a convenient alias for Root().New
func New(ctx ...interface{}) Logger {
	return root.New(ctx...)
}

// Root returns the root logger
func Root() Logger {
	return root
}

// The following functions bypass the exported logger methods (logger.Debug,
// etc.) to keep the call depth the same for all paths to logger.write so
// runtime.Caller(2) always refers to the call site in client code.

// Debug is a convenient alias for Root().Debug
func Debug(format string, a ...interface{}) {
	root.Debug(format, a...)
}

// Extra is a convenient alias for Root().Extra
func Extra(format string, a ...interface{}) {
	root.Extra(format, a...)
}

// Info is a convenient alias for Root().Info
func Info(format string, a ...interface{}) {
	root.Info(format, a...)
}

// Warn is a convenient alias for Root().Warn
func Warn(format string, a ...interface{}) {
	root.Warn(format, a...)
}

// Error is a convenient alias for Root().Error
func Error(format string, a ...interface{}) {
	root.Error(format, a...)
}

// Crit is a convenient alias for Root().Crit
func Crit(format string, a ...interface{}) {
	root.Crit(format, a...)
}
