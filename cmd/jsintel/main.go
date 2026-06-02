package main

import (
	"flag"
	"fmt"

	"github.com/Rohith-cyber2617/JS-Intel/internal/banner"
	helpmenu "github.com/Rohith-cyber2617/JS-Intel/internal/help"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display help menu")
	helpShort := flag.Bool("h", false, "Display help menu")

	flag.Parse()

	if *helpFlag || *helpShort {
		fmt.Println(banner.Banner)
		fmt.Println(helpmenu.HelpMenu)
		return
	}

	fmt.Println(banner.Banner)
}
