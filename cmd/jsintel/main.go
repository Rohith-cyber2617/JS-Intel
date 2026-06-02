package main

import (
	"fmt"
	"os"

	"github.com/Rohith-cyber2617/JS-Intel/internal/banner"
	"github.com/Rohith-cyber2617/JS-Intel/internal/output"
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

	report, err := scanner.Run(opts.URL)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		os.Exit(1)
	}

	if opts.Output != "" {

		err := output.Save(
			opts.Output,
			report.Get(),
		)

		if err != nil {
			fmt.Printf("[ERROR] Failed saving report: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("[INFO] Report saved: %s\n", opts.Output)
	}
}
