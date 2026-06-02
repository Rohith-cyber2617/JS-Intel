package finder

import (
	"regexp"

	"github.com/Rohith-cyber2617/JS-Intel/internal/models"
)

var patterns = map[string]*regexp.Regexp{

	"AWS_ACCESS_KEY": regexp.MustCompile(`AKIA[0-9A-Z]{16}`),

	"GOOGLE_API_KEY": regexp.MustCompile(`AIza[0-9A-Za-z\-_]{35}`),

	"GITHUB_TOKEN": regexp.MustCompile(`ghp_[A-Za-z0-9]{36}`),

	"JWT": regexp.MustCompile(`eyJ[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+`),

	"BEARER_TOKEN": regexp.MustCompile(`Bearer\s+[A-Za-z0-9\-\._]+`),

	"PRIVATE_KEY": regexp.MustCompile(`-----BEGIN PRIVATE KEY-----`),

	"MONGODB_URI": regexp.MustCompile(`mongodb://[^\s"'<>]+`),

	"POSTGRES_URI": regexp.MustCompile(`postgres://[^\s"'<>]+`),

	"MYSQL_URI": regexp.MustCompile(`mysql://[^\s"'<>]+`),
}

func FindSecrets(content string) []models.Finding {

	var findings []models.Finding

	for findingType, pattern := range patterns {

		matches := pattern.FindAllString(content, -1)

		for _, match := range matches {

			findings = append(findings, models.Finding{
				Type:       findingType,
				Value:      match,
				Confidence: 9.0,
				Verified:   false,
			})
		}
	}

	return findings
}
