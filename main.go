package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

var quietProcess bool

// flags
type Config struct {
	quiet   bool
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.quiet, "q", false, "")
	flag.BoolVar(&cfg.quiet, "quiet", false, "suppress non-error output (implies --noprogress)")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTION] exec "task"`)
		fmt.Fprintln(os.Stderr, "       "+name+` [OPTION] COMMAND timer-name

 COMMANDS:

  start: start named timer
  read:  read timer (elapsed time)
  stop:  read and then clear timer
  list:  list timers
  clear: clear named timer, remove from registry
  exec:  execute task and print elapsed time

OPTIONS:

  -q, --quiet
        hide process output
  -?, --help
        display this help message
  -v, --version
        print version and exit

EXAMPLES:
`)
		fmt.Fprintln(os.Stderr, "  $ "+name+` start t1`)
		fmt.Fprintln(os.Stderr, "  $ "+name+` read t1
    Elapsed time (t1): 5.9200225s`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	cmd := flag.Arg(0)
	name := flag.Arg(1)

	// Note: if flag not specified in first position, it will be ignored.
	// Example: ./timer start azerty -C ls --> flag '-C' will be ignored
	if flag.NArg() == 0 || flag.NArg() > 2 {
		flag.Usage()
		os.Exit(1)
	}
	if (cmd == "start" || cmd == "stop") && name == "" {
		fmt.Fprint(os.Stderr, "Please specify the name of the timer.\n")
		os.Exit(1)
	}
	if cmd == "exec" && name == "" {
		fmt.Fprint(os.Stderr, "Please specify a task to execute.\n")
		os.Exit(1)
	}

	// TODO: refactor
	quietProcess = cfg.quiet

	switch cmd {
	case "start":
		timer.start(name)
	case "read":
		timer.read(name)
	case "stop":
		timer.stop(name)
	case "list":
		timer.list()
	case "clear":
		timer.clear(name)
	case "exec":
		timer.exec(name)
	default:
		flag.Usage()
	}
}
