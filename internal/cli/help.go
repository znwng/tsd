package cli

import "fmt"

var helpText = `Usage: tsd [flags] <config>

Flags:
    --show <config_file>
        desc: show the passed config
        usage: tsd --show <config_file>

    --attach <session>
        desc: attach to a session after setup
        usage: tsd --attach <session-name> <config_file>

    --single <config_file>
        desc: attach to the session after creating.
              Works only when a single session is declared.
        usage: tsd --single <config_file>

    --help
        desc: print help
        usage: tsd --help`

func PrintHelp() {
	fmt.Println(helpText)
}
