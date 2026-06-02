package parser

import (
	"flag"

	"github.com/Rohith-cyber2617/JS-Intel/internal/config"
)

func ParseFlags() *config.Options {
	opts := &config.Options{}

	flag.StringVar(&opts.URL, "u", "", "Target URL")
	flag.StringVar(&opts.URL, "url", "", "Target URL")

	flag.StringVar(&opts.List, "l", "", "Target list")
	flag.StringVar(&opts.List, "list", "", "Target list")

	flag.StringVar(&opts.Output, "o", "", "Output file")
	flag.StringVar(&opts.Output, "output", "", "Output file")

	flag.IntVar(&opts.Threads, "t", 10, "Threads")
	flag.IntVar(&opts.Threads, "threads", 10, "Threads")

	flag.IntVar(&opts.Depth, "depth", 1, "Depth")

	flag.BoolVar(&opts.Endpoints, "ep", false, "Extract endpoints")
	flag.BoolVar(&opts.FoundOnly, "fo", false, "Interesting findings only")

	flag.BoolVar(&opts.Verify, "verify", false, "Verify findings")
	flag.BoolVar(&opts.RandomAgent, "random-agent", false, "Random User-Agent")

	flag.BoolVar(&opts.Silent, "silent", false, "Silent mode")

	flag.BoolVar(&opts.Update, "update", false, "Update")
	flag.BoolVar(&opts.Update, "up", false, "Update")

	flag.Parse()

	return opts
}
