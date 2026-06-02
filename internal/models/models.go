package models

type JSFile struct {
	URL         string    `json:"url" xml:"url"`
	Findings    []Finding `json:"findings,omitempty" xml:"findings>finding,omitempty"`
	RiskScore   float64   `json:"risk_score" xml:"risk_score"`
	Verified    bool      `json:"verified" xml:"verified"`
}

type Finding struct {
	Type       string  `json:"type" xml:"type"`
	Value      string  `json:"value" xml:"value"`
	Confidence float64 `json:"confidence" xml:"confidence"`
	Verified   bool    `json:"verified" xml:"verified"`
}

type Endpoint struct {
	URL      string `json:"url" xml:"url"`
	Status   int    `json:"status" xml:"status"`
	Internal bool   `json:"internal" xml:"internal"`
}

type ScanResult struct {
	Target    string     `json:"target" xml:"target"`
	JSFiles   []JSFile   `json:"js_files" xml:"js_files>js_file"`
	Endpoints []Endpoint `json:"endpoints" xml:"endpoints>endpoint"`
}
