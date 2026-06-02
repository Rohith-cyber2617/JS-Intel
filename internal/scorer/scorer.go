package scorer

func CalculateConfidence(findingType string) float64 {

	switch findingType {

	case "PRIVATE_KEY":
		return 10.0

	case "AWS_ACCESS_KEY":
		return 9.8

	case "GITHUB_TOKEN":
		return 9.5

	case "GOOGLE_API_KEY":
		return 9.2

	case "MONGODB_URI":
		return 9.0

	case "POSTGRES_URI":
		return 9.0

	case "MYSQL_URI":
		return 9.0

	case "BEARER_TOKEN":
		return 8.8

	case "JWT":
		return 8.5

	default:
		return 5.0
	}
}

func CalculateRisk(totalFindings int) float64 {

	switch {

	case totalFindings >= 20:
		return 10.0

	case totalFindings >= 15:
		return 9.0

	case totalFindings >= 10:
		return 8.0

	case totalFindings >= 5:
		return 7.0

	case totalFindings >= 1:
		return 5.0

	default:
		return 0.0
	}
}
