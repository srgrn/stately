package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var versionFlag bool

const branch = "v0.1.0"

func Version() string {
	return branch
}

type source struct {
	Target     string
	Url        string
	Branch     string
	SourceType *vcs
}

type Config struct {
	Name       string
	BaseBranch string
	Sources    []source
}

func (s *source) set_type() {
	s.SourceType = vcsGit
}

type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'stately help' output.
	Short string

	// Long is the long message shown in the 'stately help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

//var configFile string

func main() {
	flag.BoolVar(&versionFlag, "version", false, "Show the version")
	flag.BoolVar(&versionFlag, "v", false, "Show the version")

	flag.Parse()
	if versionFlag {
		fmt.Println("Stately ", Version())
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(0)
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if !cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				fmt.Println(args[1:])
				args = cmd.Flag.Args()
				fmt.Println(args)
			}
			cmd.Run(cmd, args)
			os.Exit(0)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "stately: unknown subcommand %q\n", args[0])
	os.Exit(2)

}

func usage() {
	var usage_string = `
Stately is a utility to assist with handling workspaces containing directories from several source repositories.
it contains the following commands:
get - clones/checkout all code according to configuration file.
update - updates an already existing directory with changes from scm.
freeze - create a configuration file from a given directory.
change_branch - change the branch of all or specific one of the project and updates the configuration file.
	`
	fmt.Println(usage_string)
}

var commands = []*Command{
	cmdGet,
	cmdUpdate,
	cmdFreeze,
	//cmdChangeBranch,
}
