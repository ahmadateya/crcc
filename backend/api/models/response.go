package models

type ScanDataResponse struct {
	Results    []ScanResult `json:"results"`
	Compliance int          `json:"compliance"`
}

type ScanResult struct {
	Title   string `json:"title"`
	Passed  bool   `json:"passed"`
	Details string `json:"details"`
}
