// Package dns implements the functionality exposed via the CLI.
package dns

import (
	"fmt"
	"io"
	"net"
	"time"

	"go.followtheprocess.codes/log"
)

// App represents the dns command line program.
type App struct {
	stdout io.Writer   // Normal program output is written here
	stderr io.Writer   // Logs and errors are written here
	logger *log.Logger // The application logger
}

// New returns a new [App].
func New(stdout, stderr io.Writer, debug bool) App {
	level := log.LevelInfo
	if debug {
		level = log.LevelDebug
	}
	logger := log.New(stderr, log.Prefix("dns"), log.WithLevel(level))

	return App{
		stdout: stdout,
		stderr: stderr,
		logger: logger,
	}
}

// Options represents the flags passed to dns.
type Options struct {
	RecordType string        // The --type flag, one of 'A', 'AAAA' etc.
	Server     net.IP        // The --server flag
	Timeout    time.Duration // The maximum time to wait for a DNS response
	All        bool          // Request all records, overrides --type
	Debug      bool          // Emit debug logs
}

// Run executes the DNS query.
func (a App) Run(target string, options Options) error {
	a.logger.Debug("Querying target", "target", target)
	a.logger.Debug("Options", "record-type", options.RecordType, "server", options.Server, "timeout", options.Timeout, "all", options.All)
	fmt.Fprintln(a.stdout, "Hello from DNS!")
	return nil
}
