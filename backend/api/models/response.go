package models

type ScanDataResponse struct {
	Results []ScanResult `json:"results"`
}

type ScanResult struct {
	Title   string `json:"title"`
	Passed  bool   `json:"passed"`
	Details string `json:"details"`
}
