package scanner

import (
	"fmt"

	"github.com/Rohith-cyber2617/JS-Intel/internal/colors"
	"github.com/Rohith-cyber2617/JS-Intel/internal/crawler"
	"github.com/Rohith-cyber2617/JS-Intel/internal/depth"
	"github.com/Rohith-cyber2617/JS-Intel/internal/downloader"
	"github.com/Rohith-cyber2617/JS-Intel/internal/extractor"
	"github.com/Rohith-cyber2617/JS-Intel/internal/finder"
	"github.com/Rohith-cyber2617/JS-Intel/internal/models"
	"github.com/Rohith-cyber2617/JS-Intel/internal/report"
	"github.com/Rohith-cyber2617/JS-Intel/internal/scorer"
	"github.com/Rohith-cyber2617/JS-Intel/internal/stats"
)

func Run(target string, scanDepth int) (*report.Report, error) {

	fmt.Println(colors.GreenText("[INFO] Discovering JavaScript Files"))

	jsFiles, err := crawler.DiscoverJS(target)
	if err != nil {
		return nil, err
	}

	rep := report.New()
	rep.SetTarget(target)

	stats.ScanStats.TargetsScanned++

	for _, js := range jsFiles {

		stats.ScanStats.JSFilesFound++

		if scanDepth == 1 {
			fmt.Println(colors.GreenText("[JS] " + js))

			rep.AddJSFile(models.JSFile{
				URL: js,
			})

			continue
		}

		content, err := downloader.Download(js)
		if err != nil {
			continue
		}

		var findings []models.Finding
		var endpoints []string

		if depth.AllowEndpoints(scanDepth) {
			endpoints = extractor.ExtractEndpoints(content)
		}

		if depth.AllowSecrets(scanDepth) {
			findings = finder.FindSecrets(content)
		}

		riskScore := scorer.CalculateRisk(len(findings))

		jsFile := models.JSFile{
			URL:       js,
			Findings:  findings,
			RiskScore: riskScore,
			Verified:  false,
		}

		rep.AddJSFile(jsFile)

		if len(findings) > 0 {

			fmt.Println(
				colors.YellowText("[JS] " + js),
			)

			for _, finding := range findings {

				stats.ScanStats.SecretsFound++

				fmt.Printf(
					"%s %s %s\n",
					colors.GreenText(fmt.Sprintf("[%.1f/10]", finding.Confidence)),
					colors.BlueText("("+finding.Type+")"),
					finding.Value,
				)
			}

		} else {

			fmt.Println(
				colors.GreenText("[JS] " + js),
			)
		}

		for _, endpoint := range endpoints {

			stats.ScanStats.EndpointsFound++

			rep.AddEndpoint(models.Endpoint{
				URL: endpoint,
			})

			fmt.Printf(
				"%s %s\n",
				colors.BlueText("[ENDPOINT]"),
				endpoint,
			)
		}
	}

	PrintStatistics()

	return rep, nil
}

func PrintStatistics() {

	fmt.Println()

	fmt.Println(colors.CyanText("═══════════════════════════════"))
	fmt.Println(colors.CyanText("          Statistics"))
	fmt.Println(colors.CyanText("═══════════════════════════════"))

	fmt.Printf("Targets Scanned : %d\n", stats.ScanStats.TargetsScanned)
	fmt.Printf("JS Files Found  : %d\n", stats.ScanStats.JSFilesFound)
	fmt.Printf("Endpoints Found : %d\n", stats.ScanStats.EndpointsFound)
	fmt.Printf("Secrets Found   : %d\n", stats.ScanStats.SecretsFound)
	fmt.Printf("GraphQL Found   : %d\n", stats.ScanStats.GraphQLFound)
}
