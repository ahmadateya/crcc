package models

type ScanDataResponse struct {
	Results    []ScanResult `json:"results"`
	Compliance float32      `json:"compliance"`
}

type ScanResult struct {
	Title   string `json:"title"`
	Passed  bool   `json:"passed"`
	Details string `json:"details"`
}
