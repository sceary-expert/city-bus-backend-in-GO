package models

type Stoppage struct {
	StoppageName string       `json:"stoppage-name"`
	Destination  []BusPreview `json:"bus-preview"`
}
