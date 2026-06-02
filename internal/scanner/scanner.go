package scanner

import (
	"fmt"

	"github.com/Rohith-cyber2617/JS-Intel/internal/colors"
	"github.com/Rohith-cyber2617/JS-Intel/internal/crawler"
	"github.com/Rohith-cyber2617/JS-Intel/internal/stats"
)

func Run(target string) ([]string, error) {

	fmt.Println(colors.GreenText("[INFO] Discovering JavaScript Files"))

	jsFiles, err := crawler.DiscoverJS(target)
	if err != nil {
		return nil, err
	}

	for _, js := range jsFiles {
		fmt.Println(colors.GreenText("[JS] " + js))
		stats.ScanStats.JSFilesFound++
	}

	return jsFiles, nil
}
