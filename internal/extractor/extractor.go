package extractor

import "regexp"

var endpointPatterns = []*regexp.Regexp{

	regexp.MustCompile(`https?://[^\s"'<>]+`),

	regexp.MustCompile(`["']\/api\/[^"'<> ]+["']`),

	regexp.MustCompile(`["']\/graphql["']`),

	regexp.MustCompile(`fetch\s*\(\s*["'][^"']+["']`),

	regexp.MustCompile(`axios\.(get|post|put|delete)\s*\(\s*["'][^"']+["']`),

	regexp.MustCompile(`\.ajax\s*\(\s*\{`),

	regexp.MustCompile(`XMLHttpRequest`),
}

func ExtractEndpoints(content string) []string {

	results := make(map[string]bool)

	for _, pattern := range endpointPatterns {

		matches := pattern.FindAllString(content, -1)

		for _, match := range matches {
			results[match] = true
		}
	}

	var endpoints []string

	for endpoint := range results {
		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}
