// Package cmd implements the dns CLI.
package cmd

import (
	"net"
	"time"

	"go.followtheprocess.codes/cli"
	"go.followtheprocess.codes/dns/internal/dns"
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

// Build returns the root dns CLI command.
func Build() (*cli.Command, error) {
	var options dns.Options
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
		cli.Flag(&options.Debug, "debug", cli.NoShortHand, false, "Emit debug logs"),
		cli.Run(func(cmd *cli.Command, args []string) error {
			app := dns.New(cmd.Stdout(), cmd.Stderr(), options.Debug)
			return app.Run(cmd.Arg("target"), options)
		}),
	)
}
