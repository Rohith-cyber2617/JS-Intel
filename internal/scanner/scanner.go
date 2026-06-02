package scanner

import (
	"fmt"

	"github.com/Rohith-cyber2617/JS-Intel/internal/colors"
	"github.com/Rohith-cyber2617/JS-Intel/internal/crawler"
	"github.com/Rohith-cyber2617/JS-Intel/internal/downloader"
	"github.com/Rohith-cyber2617/JS-Intel/internal/extractor"
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

		content, err := downloader.Download(js)
		if err != nil {
			continue
		}

		endpoints := extractor.ExtractEndpoints(content)

		for _, endpoint := range endpoints {

			stats.ScanStats.EndpointsFound++

			fmt.Printf(
				"%s %s\n",
				colors.BlueText("[ENDPOINT]"),
				endpoint,
			)
		}
	}

	return jsFiles, nil
}
