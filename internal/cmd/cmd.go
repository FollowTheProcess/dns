// Package cmd implements the dns CLI.
package cmd

import (
	"fmt"
	"net"
	"time"

	"go.followtheprocess.codes/cli"
)

var (
	version = "dev"
	commit  = ""
	date    = ""
)

const defaultTimeout = 3 * time.Second

const longAbout = `
dns is a flexible command line tool for performing DNS queries.

It aims to be broadly equivalent on a feature level with dig but with a more
modern and intuitive command line and UX.

Like dig, if not told otherwise, dns looks in /etc/resolve.conf (if it exists) for name servers and will
try each of them in turn, returning the first successful response.

You may provide the address of any given name server with the '--server' flag.

The query, by default, will ask for A records from the name server. This may be set by the '--type' flag
or all records requested with the '--all' flag.
`

// Options represents the flags passed to dns.
type Options struct {
	RecordType string        // The --type flag, one of 'A', 'AAAA' etc.
	Server     net.IP        // The --server flag
	Timeout    time.Duration // The maximum time to wait for a DNS response
	All        bool          // Request all records, overrides --type
}

// Build returns the root dns CLI command.
func Build() (*cli.Command, error) {
	var options Options
	return cli.New(
		"dns",
		cli.Short("DNS query cli tool, a modern alternative to dig"),
		cli.Long(longAbout),
		cli.Version(version),
		cli.Commit(commit),
		cli.BuildDate(date),
		cli.RequiredArg("target", "The target of the DNS query, typically a domain name"),
		cli.Flag(&options.RecordType, "type", cli.NoShortHand, "A", "The type of record to query for"),
		cli.Flag(&options.Server, "server", cli.NoShortHand, net.ParseIP("1.1.1.1"), "The address of the name server"),
		cli.Flag(&options.Timeout, "timeout", cli.NoShortHand, defaultTimeout, "The maximum time to wait for a DNS response"),
		cli.Run(func(cmd *cli.Command, args []string) error {
			fmt.Println("Hello")
			return nil
		}),
	)
}
