package report

import "github.com/Rohith-cyber2617/JS-Intel/internal/models"

type Report struct {
	Result models.ScanResult
}

func New() *Report {
	return &Report{
		Result: models.ScanResult{},
	}
}

func (r *Report) SetTarget(target string) {
	r.Result.Target = target
}

func (r *Report) AddJSFile(js models.JSFile) {
	r.Result.JSFiles = append(r.Result.JSFiles, js)
}

func (r *Report) AddEndpoint(endpoint models.Endpoint) {
	r.Result.Endpoints = append(r.Result.Endpoints, endpoint)
}

func (r *Report) Get() models.ScanResult {
	return r.Result
}
