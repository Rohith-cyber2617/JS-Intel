package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Rohith-cyber2617/JS-Intel/internal/banner"
	helpmenu "github.com/Rohith-cyber2617/JS-Intel/internal/help"
	"github.com/Rohith-cyber2617/JS-Intel/internal/parser"
	"github.com/Rohith-cyber2617/JS-Intel/internal/validator"
)

func main() {

	helpFlag := flag.Bool("help", false, "Display help menu")
	helpShort := flag.Bool("h", false, "Display help menu")

	opts := parser.ParseFlags()

	if *helpFlag || *helpShort {
		fmt.Println(banner.Banner)
		fmt.Println(helpmenu.HelpMenu)
		return
	}

	fmt.Println(banner.Banner)

	if err := validator.Validate(opts); err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}

	fmt.Println("[INFO] Configuration Loaded")

	if opts.URL != "" {
		fmt.Printf("[INFO] Target URL : %s\n", opts.URL)
	}

	if opts.List != "" {
		fmt.Printf("[INFO] Target List : %s\n", opts.List)
	}

	fmt.Printf("[INFO] Threads    : %d\n", opts.Threads)
	fmt.Printf("[INFO] Depth      : %d\n", opts.Depth)

	if opts.Verify {
		fmt.Println("[INFO] Verification Enabled")
	}

	if opts.RandomAgent {
		fmt.Println("[INFO] Random User-Agent Enabled")
	}
}
