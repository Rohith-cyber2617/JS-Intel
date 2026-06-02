package stats

type Statistics struct {
	TargetsScanned  int
	JSFilesFound    int
	EndpointsFound  int
	ParametersFound int
	SecretsFound    int
	GraphQLFound    int
	AdminRoutes     int
	InternalAPIs    int
	VerifiedFinds   int
	ScanDuration    string
}

var ScanStats = Statistics{}
