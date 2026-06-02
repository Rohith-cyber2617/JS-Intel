package finder

import (
	"regexp"

	"github.com/Rohith-cyber2617/JS-Intel/internal/models"
	"github.com/Rohith-cyber2617/JS-Intel/internal/scorer"
)

var patterns = map[string]*regexp.Regexp{

	// Cloud Keys

	"AWS_ACCESS_KEY": regexp.MustCompile(`AKIA[0-9A-Z]{16}`),

	"GOOGLE_API_KEY": regexp.MustCompile(`AIza[0-9A-Za-z\-_]{35}`),

	// Git

	"GITHUB_TOKEN": regexp.MustCompile(`ghp_[A-Za-z0-9]{36}`),

	"GITLAB_TOKEN": regexp.MustCompile(`glpat-[A-Za-z0-9\-_]+`),

	// Tokens

	"JWT": regexp.MustCompile(`eyJ[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+`),

	"BEARER_TOKEN": regexp.MustCompile(`Bearer\s+[A-Za-z0-9\-\._]+`),

	// Databases

	"MONGODB_URI": regexp.MustCompile(`mongodb://[^\s"'<>]+`),

	"POSTGRES_URI": regexp.MustCompile(`postgres://[^\s"'<>]+`),

	"MYSQL_URI": regexp.MustCompile(`mysql://[^\s"'<>]+`),

	// Keys

	"PRIVATE_KEY": regexp.MustCompile(`-----BEGIN PRIVATE KEY-----`),

	// Firebase

	"FIREBASE": regexp.MustCompile(`firebase`),

	// Stripe

	"STRIPE_SECRET": regexp.MustCompile(`sk_live_[A-Za-z0-9]+`),

	// Twilio

	"TWILIO": regexp.MustCompile(`twilio`),

	// SendGrid

	"SENDGRID": regexp.MustCompile(`sendgrid`),

	// Admin

	"ADMIN_ROUTE": regexp.MustCompile(`admin`),

	"SUPERADMIN": regexp.MustCompile(`superadmin`),

	"ADMIN_PANEL": regexp.MustCompile(`adminPanel`),

	// Internal

	"INTERNAL_API": regexp.MustCompile(`internal_api`),

	"PRIVATE_API": regexp.MustCompile(`private_api`),

	// GraphQL

	"GRAPHQL": regexp.MustCompile(`graphql`),

	"INTROSPECTION": regexp.MustCompile(`introspection`),

	// Feature Flags

	"FEATURE_FLAG": regexp.MustCompile(`featureFlag`),

	"ENABLE_ADMIN": regexp.MustCompile(`enableAdmin`),

	// Debug

	"DEBUG_MODE": regexp.MustCompile(`debugMode`),

	"DEV_ONLY": regexp.MustCompile(`devOnly`),

	// Comments

	"TODO": regexp.MustCompile(`TODO:`),

	"FIXME": regexp.MustCompile(`FIXME:`),

	"HACK": regexp.MustCompile(`HACK:`),
}

func FindSecrets(content string) []models.Finding {

	var findings []models.Finding

	for findingType, pattern := range patterns {

		matches := pattern.FindAllString(content, -1)

		for _, match := range matches {

			findings = append(findings, models.Finding{
				Type:       findingType,
				Value:      match,
				Confidence: scorer.CalculateConfidence(findingType),
				Verified:   false,
			})
		}
	}

	return findings
}
