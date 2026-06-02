package main

import (
	"fmt"
	"os"

	"github.com/Rohith-cyber2617/JS-Intel/internal/banner"
	"github.com/Rohith-cyber2617/JS-Intel/internal/parser"
	"github.com/Rohith-cyber2617/JS-Intel/internal/scanner"
	"github.com/Rohith-cyber2617/JS-Intel/internal/validator"
)

func main() {

	opts := parser.ParseFlags()

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

	fmt.Printf("[INFO] Threads : %d\n", opts.Threads)
	fmt.Printf("[INFO] Depth   : %d\n", opts.Depth)

	if opts.Verify {
		fmt.Println("[INFO] Verification Enabled")
	}

	if opts.RandomAgent {
		fmt.Println("[INFO] Random User-Agent Enabled")
	}

	if opts.URL != "" {
		_, err := scanner.Run(opts.URL)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err)
			os.Exit(1)
		}
	}
}
